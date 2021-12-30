.DEFAULT_GOAL := build

ARTIFACT = libfal.so
BUILDMODE = c-shared
LDFLAGS =

build:
	go build $(LDFLAGS) -o $(ARTIFACT) -buildmode=$(BUILDMODE) .

release: LDFLAGS = -ldflags "-w -s"
release: build

binary: ARTIFACT = fal
binary: BUILDMODE = exe
binary: build
