FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
ENV http_proxy="http://46.105.221.241:33080" \
        https_proxy="http://46.105.221.241:33080"

ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

CMD ["./consignment-cli"]
