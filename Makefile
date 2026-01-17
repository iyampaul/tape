APP_NAME = tape
BINARY = $(APP_NAME).bin
SRC = ./tape/main.go

.PHONY: all build run clean docker docker-build docker-up docker-down

all: build

build:
	go build -o $(BINARY) $(SRC)

run:
	sudo ./$(BINARY)

clean:
	rm -f $(BINARY)

docker: docker-build docker-up

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
