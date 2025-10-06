.PHONY: run build

build:
	go build -o challenge_tracker_service

run: build
	./challenge_tracker_service
