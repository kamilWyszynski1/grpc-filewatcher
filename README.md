# FILEWATCHER GRPC
Repository created in order to learn some of gRPC and protobufs.

## Flow
1. Server starts with hardcoded filepath. 
2. Manager create goroutines for watching over changes in that file. Additionally starts repository listener for saving these changes.
3. gRPC allows to retrive changes and order new watcher-flow.

## Technologies used
* [boltdb](https://github.com/boltdb/bolt)
* [grpc-go](https://grpc.io/docs/languages/go/quickstart/)
* [protobuf](https://developers.google.com/protocol-buffers)
* [fsnotify](https://github.com/fsnotify/fsnotify)