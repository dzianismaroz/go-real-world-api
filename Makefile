clean:
	go clean
build:
	go build ./cmd/main.go
run:
	go run ./cmd/main.go
tests:
	go test  ./...
vet:
	go vet


