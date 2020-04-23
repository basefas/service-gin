.PHONY: run build

run:
	go run ./cmd/app/main.go
build:
	go build -o ./build/service-gin -v ./cmd/app/main.go