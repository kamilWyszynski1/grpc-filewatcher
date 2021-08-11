package main

import (
	"context"
	"fmt"
	"log"

	watcher2 "github.com/kamilWyszynski1/filewatcher-grpc/internal/watcher"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := watcher2.NewWatcherServiceClient(cc)
	request := &watcher2.Empty{}

	resp, _ := client.GetLastChange(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)
}
