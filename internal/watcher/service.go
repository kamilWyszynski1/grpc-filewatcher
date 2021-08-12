package watcher

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kamilWyszynski1/filewatcher-grpc/internal/pb"

	"github.com/fsnotify/fsnotify"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Manager struct {
	pb.UnimplementedWatcherServiceServer
	log                *log.Logger
	done               map[string]chan struct{}
	repo               Repository
	changesChannels    map[*pb.StartWatchingRequest]chan *pb.Change
	fileWatcherClosers map[*pb.StartWatchingRequest]func() error
}

func NewManager(repo Repository, options ...Option) *Manager {
	s := &Manager{
		repo:               repo,
		fileWatcherClosers: map[*pb.StartWatchingRequest]func() error{},
		done:               map[string]chan struct{}{},
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

type Option func(m *Manager)

func WithLogger(log *log.Logger) Option {
	return func(m *Manager) {
		m.log = log
	}
}

func WithListener(l Listener) Option {
	return func(m *Manager) {
		m.addListener(l)
	}
}

func (m *Manager) addListener(l Listener) {
	if m.changesChannels == nil {
		m.changesChannels = make(map[*pb.StartWatchingRequest]chan *pb.Change)
	}
	m.changesChannels[l.StartWatchingRequest] = l.Channel
}

func (m Manager) Close() error {
	for _, d := range m.done {
		d <- struct{}{}
	}
	for fName, c := range m.fileWatcherClosers {
		if err := c(); err != nil {
			return fmt.Errorf("failed to close %s watcher", fName)
		}
	}
	return nil
}

// Watch main function that starts whole watching over file.
// Method starts goroutine that watch files and goroutines that tracks changes in repo.
func (m *Manager) Watch() error {
	for req, channel := range m.changesChannels {
		if err := m.startWatchingFlow(req, channel); err != nil {
			return err
		}
	}
	return nil
}

// startWatchingFlow starts repo listener and watcher goroutine.
func (m *Manager) startWatchingFlow(req *pb.StartWatchingRequest, channel chan *pb.Change) error {
	if m.isFileAlreadyWatched(req.FilePath) {
		m.logf("file %s already watched, skipping\n", req.FileAlias)
		return nil
	}
	fWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	if err := m.repo.StartListening(Listener{
		StartWatchingRequest: req,
		Channel:              channel,
	}); err != nil {
		return fmt.Errorf("failed to start listening on repo, %w", err)
	}
	m.fileWatcherClosers[req] = fWatcher.Close
	stopChannel := make(chan struct{})
	m.done[req.FileAlias] = stopChannel

	m.logn(fmt.Sprintf("starting watching over %s file", req.FileAlias))
	go m.watch(req.FileAlias, fWatcher, channel, stopChannel)
	return fWatcher.Add(req.FilePath)
}

func (m Manager) isFileAlreadyWatched(filepath string) bool {
	for k := range m.fileWatcherClosers {
		if k.FilePath == filepath {
			return true
		}
	}
	return false
}

func (m Manager) watch(fileAlias string, fileWatcher *fsnotify.Watcher, changes chan *pb.Change, done chan struct{}) {
	for {
		select {
		case event, ok := <-fileWatcher.Events:
			if !ok {
				continue
			}
			ch := &pb.Change{
				FileAlias: fileAlias,
				EventName: event.String(),
				Timestamp: timestamppb.Now(),
			}
			changes <- ch
		case err, ok := <-fileWatcher.Errors:
			if !ok {
				continue
			}
			m.logn("error:", err)

		case <-done:
			m.logn("closing watcher")
			return
		}
	}
}

func (m Manager) logf(format string, args ...interface{}) {
	if m.log != nil {
		m.log.Printf(format, args...)
	}
}

func (m Manager) logn(args ...interface{}) {
	if m.log != nil {
		m.log.Println(args...)
	}
}

// GetChanges returns all file changes.
func (m Manager) GetChanges(context.Context, *pb.GetChangesRequest) (*pb.GetChangesResponse, error) {
	changes, err := m.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.GetChangesResponse{Changes: changes}, nil
}

// StartWatching order manager to create new watch-flow for file.
func (m *Manager) StartWatching(_ context.Context, req *pb.StartWatchingRequest) (*empty.Empty, error) {
	m.logn(fmt.Sprintf("starting watching over: %s", req.FilePath))
	ch := make(chan *pb.Change)
	m.addListener(Listener{
		Channel:              ch,
		StartWatchingRequest: req,
	})
	return &empty.Empty{}, m.startWatchingFlow(req, ch)
}
