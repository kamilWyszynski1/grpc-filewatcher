package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kamilWyszynski1/filewatcher-grpc/internal/pb"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewWatcherServiceClient(cc)
	request := &pb.GetChangesRequest{FileAlias: ""}

	resp, _ := client.GetChanges(context.Background(), request)
	fmt.Println(len(resp.Changes))
	for _, r := range resp.Changes {
		fmt.Printf("%s :: alias: %s, event: %s\n", r.Timestamp.AsTime().Format("20060102 15:04:05"), r.FileAlias, r.EventName)
	}

	client.StartWatching(context.Background(), &pb.StartWatchingRequest{
		FileAlias: "test2",
		FilePath:  "internal/testdata/testfile2.txt",
	})
}
