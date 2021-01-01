image = goff:dev

.PHONY: build
build:
	docker build -t $(image) $(PWD)

.PHONY: run
run:
	docker run -it -v $(PWD):/app/src/github.com/miniav/goff --rm $(image) sh

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: test
test:
	go test ./...
