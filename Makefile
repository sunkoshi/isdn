test:
	CGO_ENABLED=0 go test -v -cover ./...
dev:
	go run *.go
demo:
	rm -rf test_codes/cppm/cppm.zip
	cd ./test_codes/cppm/ && zip -r cppm.zip *
	go run *.go
	