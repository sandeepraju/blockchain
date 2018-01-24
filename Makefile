.PHONY: all clean run build setup

all:
	clean

clean:
	rm -rf ./blockchain

run:
	go run main.go

build:
	go build .

setup:
	dep ensure