test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -v ./...

check: vet fmt build test

clean:
	# rm everything except pdf files in assets
	go clean
