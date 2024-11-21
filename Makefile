.PHONY: dev
dev:
	$(HOME)/go/bin/air

.PHONY: build
build:
	go build -o ./tmp/main .

.PHONY: run
run:
	./tmp/main
