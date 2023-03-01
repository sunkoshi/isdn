test:
	CGO_ENABLED=0 go test -v -cover ./...
dev:
	go run *.go
demo:
	rm -rf test_codes/jsm/jsm.zip
	cd ./test_codes/jsm/ && zip -r jsm.zip *
	go run *.go
	
full:
	go run ./cmd/function_server/*.go