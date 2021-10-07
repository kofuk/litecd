GO = go

.PHONY: all
all:
	$(GO) build

.PHONY: run
run:
	$(GO) run .

.PHONY: clean
clean:
	$(RM) litecd
