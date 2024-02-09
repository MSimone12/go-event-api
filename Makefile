run:
	go run ./src/main.go

docker:
	docker build --platform="linux/amd64" -t go-ebanx .
	docker tag go-ebanx:latest 954461865512.dkr.ecr.us-east-2.amazonaws.com/go-ebanx:latest
	docker push 954461865512.dkr.ecr.us-east-2.amazonaws.com/go-ebanx:latest

ci: prebuild build