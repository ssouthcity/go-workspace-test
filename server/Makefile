proto:
	protoc -I=proto --go_out=gen --go-grpc_out=gen proto/*.proto

clean:
	rm -rf ./gen/*

.PHONY: proto clean