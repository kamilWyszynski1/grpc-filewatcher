package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/lru"
	wtch "github.com/kamilWyszynski1/filewatcher-grpc/internal/watcher"
	"google.golang.org/grpc"
)

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	cache := lru.NewLRU(10, nil)
	svc := wtch.NewService(cache)
	err = svc.
		Logger(log.New(os.Stdout, "", log.LstdFlags)).
		Filename("../internal/testdata/testfile.txt").
		Watch()
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	wtch.RegisterWatcherServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
