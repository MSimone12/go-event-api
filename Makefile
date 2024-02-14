run: tidy
	NGROK_AUTHTOKEN=2c826go2wz0f5TsgsmqJFyDQHZf_34EGLBswNxA6C5dHjC7Fq go run ./src/main.go

docker:
	docker build --platform="linux/amd64" -t go-ebanx .
	docker tag go-ebanx:latest 954461865512.dkr.ecr.us-east-2.amazonaws.com/go-ebanx:latest
	docker push 954461865512.dkr.ecr.us-east-2.amazonaws.com/go-ebanx:latest

tidy:
	go mod tidy