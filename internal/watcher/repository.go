package watcher

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// TODO: implement observer pattern

// RepositoryListener is an API for handling async data-layer of watcher.
type RepositoryListener interface {
	// StartListening starts listening on channel and saves every single item received.
	StartListening(chan *Change) error
	StopListening()
}

// Repository handle sync
type Repository interface {
	Get(id uint64) (*Change, error)
	GetAll() ([]*Change, error)
}

type repository struct {
	db   *bolt.DB
	log  log.Logger
	done chan struct{}
}

func NewRepository(db *bolt.DB) *repository {
	return &repository{db: db}
}

const bucketName = "changes"

func (r *repository) StartListening(changes chan *Change) error {
	if err := r.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketName)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-r.done:
				return
			case change := <-changes:
				r.handleMessage(change)
			}
		}
	}()
	return nil
}

func (r *repository) handleMessage(change *Change) {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
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

func (r *repository) Get(id uint64) (*Change, error) {
	var resp *Change
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		data := b.Get(itob(int(id)))

		var ch Change
		if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&ch); err != nil {
			return err
		}
		resp = &ch
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *repository) GetAll() ([]*Change, error) {
	changes := make([]*Change, 0)
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.ForEach(func(k, data []byte) error {
			var ch Change
			fmt.Println(binary.BigEndian.Uint64(k))
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
