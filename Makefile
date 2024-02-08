run:
	go run main.go
build:
	go build -o bin/main main.go 

prebuild:
	echo "Running prebuild..."
	go mod edit -go=1.18.10
	echo "Ran go mod edit..."
	go mod tidy
	echo "Ran go mod tidy..."

postbuild:
	echo "Running prebuild..."
	go mod edit -go=1.21.6
	echo "Ran go mod edit..."

ci: prebuild build postbuild