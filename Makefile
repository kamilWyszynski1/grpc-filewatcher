gen-proto:
	protoc --go_out=internal/watcher/ --go-grpc_out=internal/watcher/ internal/watcher/watcher.proto
