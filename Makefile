.PHONY: clean
clean:
	rm go-circle-list-extract*

.PHONY: test
test:
	go test -v ./...

.PHONY: verify
verify:
	go mod download
	go mod verify

.PHONY: build
build: clean
	go build .
