package watcher

import (
	"context"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/kamilWyszynski1/filewatcher-grpc/internal/lru"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	UnimplementedWatcherServiceServer
	log               *log.Logger
	filename          string
	cache             lru.LRU
	fileWatcherCloser func() error
	done              chan struct{}
	repo              Repository
}

func NewService(cache lru.LRU, repo Repository) *Service {
	return &Service{
		cache: cache,
		repo:  repo,
	}
}

func (s *Service) Filename(filename string) *Service {
	s.filename = filename
	return s
}

func (s *Service) Logger(log *log.Logger) *Service {
	s.log = log
	return s
}

func (s Service) Close() error {
	s.done <- struct{}{}
	return s.fileWatcherCloser()
}

func (s *Service) Watch(changes chan *Change) error {
	fileWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	s.fileWatcherCloser = fileWatcher.Close

	go func() {
		for {
			select {
			case event, ok := <-fileWatcher.Events:
				if !ok {
					continue
				}
				fmt.Println(event)
				ch := &Change{
					FileName:  s.filename,
					EventName: event.String(),
					Timestamp: timestamppb.Now(),
				}
				s.cache.Add(ch)
				changes <- ch
			case err, ok := <-fileWatcher.Errors:
				if !ok {
					continue
				}
				s.logn("error:", err)

			case <-s.done:
				s.logn("closing watcher")
				return
			}
		}
	}()
	if err := fileWatcher.Add(s.filename); err != nil {
		return err
	}
	return nil
}

func (s Service) logf(format string, args ...interface{}) {
	if s.log != nil {
		s.log.Printf(format, args...)
	}
}

func (s Service) logn(args ...interface{}) {
	if s.log != nil {
		s.log.Println(args...)
	}
}

func (s Service) GetLastChange(context.Context, *Empty) (*Change, error) {
	return s.repo.Get(1)
}
