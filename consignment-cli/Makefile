build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .
RUN:
	docker run -e MICRO_REGISTRY=mdns consignment-cli