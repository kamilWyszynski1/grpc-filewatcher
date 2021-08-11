package watcher

import (
	"fmt"
	"testing"

	bolt "go.etcd.io/bbolt"
)

func Test_repository_GetAll(t *testing.T) {
	db, err := bolt.Open("../../server/bolt.db", 0666, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(NewRepository(db).GetAll())
}
