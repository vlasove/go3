.PHONY:proto
proto:
	protoc -I. --go_out=plugins=micro:. proto/user/user.proto 

.PHONY:build
build:
	sudo docker build -t userserver .

.PHONY:run
run:
	sudo docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns userserver 


.DEFAULT_GOAL := proto