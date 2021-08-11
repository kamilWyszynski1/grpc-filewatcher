package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/lru"
	pb "github.com/kamilWyszynski1/filewatcher-grpc/internal/watcher"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
)

const (
	address  = "0.0.0.0:50051"
	boltPath = "./bolt.db"
)

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	db, err := bolt.Open(boltPath, 0666, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	changes := make(chan *pb.Change)

	cache := lru.NewLRU(10, nil)
	repo := pb.NewRepository(db)
	if err := repo.StartListening(changes); err != nil {
		panic(err)
	}
	svc := pb.NewService(cache, repo)
	err = svc.
		Logger(log.New(os.Stdout, "", log.LstdFlags)).
		Filename("internal/testdata/testfile.txt").
		Watch(changes)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterWatcherServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
