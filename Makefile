NAME = nucktwillieren/session-go

.PHONY: build start push

lint: 
	golangci-lint run ./...

test: 
	go test -p 1 -v ./...
	
