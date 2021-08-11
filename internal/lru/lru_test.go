package lru_test

import (
	"fmt"
	"testing"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/lru"
)

func TestLRU(t *testing.T) {
	cache := lru.NewLRU(5, func(v interface{}) {
		fmt.Println(v)
	})

	for i := 0; i < 10; i++ {
		cache.Add(i)
	}

}
