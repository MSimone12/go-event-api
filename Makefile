run:
	go run main.go
build:
	go build main.go

prebuild:
	go mod edit -go=1.18.10
	go mod tidy

postbuild:
	go mod edit -go=1.21.6