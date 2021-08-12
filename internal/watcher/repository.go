package watcher

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/pb"

	bolt "go.etcd.io/bbolt"
)

// TODO: implement observer pattern

// Repository handle sync
type Repository interface {
	// Get returns all filename changes.
	Get(filename string) ([]*pb.Change, error)
	// GetAll returns all changes noticed.
	GetAll() ([]*pb.Change, error)

	// StartListening starts listening on channel and saves every single item received.
	StartListening(listeners Listener) error
	StopListening()
}

type repository struct {
	db   *bolt.DB
	log  *log.Logger
	done chan struct{}
}

func NewRepository(db *bolt.DB, log *log.Logger) *repository {
	return &repository{db: db, log: log}
}

func (r *repository) StartListening(listener Listener) error {
	if err := r.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(listener.FileAlias)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	r.log.Printf("starting repo listener for %s file\n", listener.FileAlias)
	go func(ch chan *pb.Change) {
		for {
			select {
			case <-r.done:
				return
			case change := <-ch:
				r.handleMessage(change)
			}
		}
	}(listener.Channel)
	return nil
}

func (r *repository) handleMessage(change *pb.Change) {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(change.FileAlias))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		r.log.Printf("saving change for %s with id %d\n", change.FileAlias, id)
		buf, err := json.Marshal(change)
		if err != nil {
			return err
		}
		return b.Put(itob(int(id)), buf)
	})
	if err != nil {
		log.Printf("failed to update, %s\n", err)
	}
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (r *repository) StopListening() {
	r.done <- struct{}{}
}

func (r *repository) Get(filename string) ([]*pb.Change, error) {
	changes := make([]*pb.Change, 0)
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(filename))
		return b.ForEach(func(_, data []byte) error {
			var ch pb.Change
			if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&ch); err != nil {
				return err
			}
			changes = append(changes, &ch)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return changes, nil
}

func (r *repository) GetAll() ([]*pb.Change, error) {
	changes := make([]*pb.Change, 0)
	err := r.db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(_ []byte, b *bolt.Bucket) error {
			return b.ForEach(func(_, data []byte) error {
				var ch pb.Change
				if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&ch); err != nil {
					return err
				}
				changes = append(changes, &ch)
				return nil
			})
		})
	})
	if err != nil {
		return nil, err
	}
	return changes, nil
}
