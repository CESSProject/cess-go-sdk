fmt:
	go fmt ./...

build:
	go build -v ./...

test:
	go test -v ./...

check: fmt build test

clean:
	# rm everything except pdf files in assets
	go clean
