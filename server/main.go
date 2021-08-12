package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/watcher"

	pb "github.com/kamilWyszynski1/filewatcher-grpc/internal/pb"

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
	fmt.Printf("Server is listening on %v ...\n", address)

	db, err := bolt.Open(boltPath, 0666, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	l := watcher.Listener{
		StartWatchingRequest: &pb.StartWatchingRequest{
			FilePath:  "internal/testdata/testfile.txt",
			FileAlias: "testfile",
		},
		Channel: make(chan *pb.Change),
	}

	repo := watcher.NewRepository(db, log.New(os.Stdout, "repo ", log.LstdFlags))
	svc := watcher.NewManager(repo,
		watcher.WithLogger(log.New(os.Stdout, "svc ", log.LstdFlags)),
		watcher.WithListener(l))
	if err := svc.Watch(); err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterWatcherServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
