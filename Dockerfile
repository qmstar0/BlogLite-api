FROM golang:1.22 AS builder
WORKDIR /BlogLite-api

COPY . .

ENV GOPROXY=https://goproxy.io/
ENV TZ=Asia/Shanghai
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -o ./build/main .

FROM alpine:latest
WORKDIR /BlogLite-api
COPY --from=builder /BlogLite-api/build/main .
CMD ["./main"]
