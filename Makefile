run: tidy
	NGROK_AUTHTOKEN=2c826go2wz0f5TsgsmqJFyDQHZf_34EGLBswNxA6C5dHjC7Fq go run ./src/main.go

tidy:
	go mod tidy

run-test:
	go test ./test/... -v