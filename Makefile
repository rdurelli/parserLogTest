build:
	go build -ldflags="-w -s" -o bin/main src/main.go

run:
	go run src/main.go