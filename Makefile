gen:
	protoc -I=proto --go_out=. proto/*.proto
clean:
	rm -rf proto-gen