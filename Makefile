GO = go

.PHONY: all
all:
	$(GO) build

.PHONY: clean
clean:
	$(RM) litecd
