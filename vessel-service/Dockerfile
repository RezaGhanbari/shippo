# vessel-service/Dockerfile
FROM golang:1.10.0 as builder

WORKDIR /go/src/shippo/vessel-service

COPY . .
ENV http_proxy="http://46.105.221.241:33080" \
        https_proxy="http://46.105.221.241:33080"

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/shippo/vessel-service/vessel-service .

CMD ["./vessel-service"]
