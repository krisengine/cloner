APP_NAME = $(shell cat APP_NAME)
VERSION = $(shell cat VERSION)

default: build

build:
	docker build -t $(APP_NAME):$(VERSION) -f Dockerfile .
	docker build -t $(APP_NAME):latest -f Dockerfile .

push: build
	docker push $(APP_NAME):$(VERSION)
	docker push $(APP_NAME):latest

run: build
	docker run -p 8000:8000 -v "${PWD}/rep:/app/rep" -v "${PWD}/config.json:/app/config.json" $(APP_NAME):latest