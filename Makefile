run:
  go run ./cmd/example -e "+ 1 2"

test:
  go test ./...

build:
  go build -o bin/app ./cmd/example
