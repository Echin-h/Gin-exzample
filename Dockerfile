FROM golang:1.21.1 as builder

WORKDIR /app

COPY . .

RUN go build -o myapp .

EXPOSE 8080

CMD ["./myapp"]
