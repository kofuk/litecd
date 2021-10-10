GO = go

.PHONY: all
all:
	$(GO) build

.PHONY: run
run:
	$(GO) run .

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	$(RM) litecd
