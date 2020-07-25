.PHONY: build
build:
	docker build -t play-image .

.PHONY: run
run:
	docker run -d -p 7531:7531 --env-file .env --restart unless-stopped --name play play-image

.DEFAULT_GOAL := run
