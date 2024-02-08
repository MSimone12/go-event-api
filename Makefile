run:
	go run ./src/main.go
build:
	go build -o main ./src

prebuild:
	rm -rf go.mod go.sum
	go mod init go-event-api
	go mod tidy

ci: prebuild build