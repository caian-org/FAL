.DEFAULT_GOAL := build

ARTIFACT = falctl
LDFLAGS =

.PHONY: build
build:
	@cd fal && go build $(LDFLAGS) -o $(ARTIFACT) .
	@mv fal/$(ARTIFACT) .

release: LDFLAGS = -ldflags "-w -s"
release: build
	@upx --best --lzma $(ARTIFACT)

format:
	gofmt -s -w fal
