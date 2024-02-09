FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
RUN go build -o main ./src

EXPOSE 8080

CMD ["./main"]