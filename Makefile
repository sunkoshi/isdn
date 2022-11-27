test:
	CGO_ENABLED=0 go test -v -cover ./...
dev:
	go run *.go
demo:
	rm -rf test_codes/pym/pym.zip
	cd ./test_codes/pym/ && zip -r pym.zip *
	go run *.go
	