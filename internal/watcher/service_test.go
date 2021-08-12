package watcher_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/lru"
	"github.com/kamilWyszynski1/filewatcher-grpc/internal/watcher"
)

func TestWatcherService(t *testing.T) {
	cache := lru.NewLRU(10, nil)
	svc := watcher.NewManager(nil)

	f, err := os.OpenFile(path.Join("..", "testdata", "testfile.txt"), os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	if err != nil {
		t.Fatal(err)
	}
	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))

	if err := svc.Filename(f.Name()).Watch(nil); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 3)
	if _, err := f.Write([]byte("siemano test")); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}
