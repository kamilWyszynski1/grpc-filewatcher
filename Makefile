gen-proto:
	protoc --go_out=internal/pb/ --go-grpc_out=internal/pb/ internal/pb/watcher.proto
