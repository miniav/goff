app_name                := goff
docker_name             := $(app_name)
docker_tag              := dev
docker_container        := $(app_name)

.PHONY: upgrade
upgrade:
	docker pull alpine:edge

.PHONY: build
build:
	docker build -t $(docker_name):$(docker_tag) .

.PHONY: run
run:
	docker-compose up --force-recreate -d

.PHONY: exec
exec:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish
