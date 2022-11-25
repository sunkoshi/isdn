test:
	CGO_ENABLED=0 go test -v -cover ./...
dev:
	go run *.go