package watcher

import "github.com/kamilWyszynski1/filewatcher-grpc/internal/pb"

type Listener struct {
	*pb.StartWatchingRequest
	Channel chan *pb.Change
}
