app_name                := goff
docker_name             := $(app_name)
docker_tag              := dev
docker_container        := $(app_name)

golang_version          := 1.12
golang_sha256           := 750a07fef8579ae4839458701f4df690e0b20b8bcce33b437e4df89c451b6f13

.PHONY: upgrade
upgrade:
	docker pull buildpack-deps:stable

.PHONY: build
build:
	docker build -t $(docker_name):$(docker_tag) .

.PHONY: run
run:
	docker-compose up --force-recreate --remove-orphans --build -d

.PHONY: exec
exec:
	docker exec -e COLUMNS="`tput cols`" -e LINES="`tput lines`" -it $(docker_container) /usr/bin/fish
