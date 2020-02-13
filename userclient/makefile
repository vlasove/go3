.PHONY:build
build:
	sudo docker build -t userclient .

.PHONY:run
run:
	sudo docker run -e MICRO_REGISTRY=mdns userclient

.DEFAULT_GOAL := build