app_name = goff

.PHONY: upgrade
upgrade:
	docker pull alpine:edge

.PHONY: build
build:
	docker-compose build $(app_name)

.PHONY: run
run:
	docker-compose up --force-recreate -d

.PHONY: exec
exec:
	docker-compose exec $(app_name) /usr/bin/fish
