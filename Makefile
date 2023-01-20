gen:
	protoc -I=proto --go_out=. proto/*.proto --go-grpc_out=. proto/*proto
clean:
	rm -rf proto-gen
test:
	go test -cover -race ./...