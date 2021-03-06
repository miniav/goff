image = goff:dev

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: image
image:
	docker build -t $(image) $(PWD)

.PHONY: docker-run
docker-run:
	docker run -it -v $(PWD):/app/src/github.com/miniav/goff --rm $(image) /bin/sh
