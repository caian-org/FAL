.DEFAULT_GOAL := build

ARTIFACT = libfal
BUILDMODE = c-shared
LDFLAGS =

EXT = dll

ifneq ($(OS),Windows_NT)
	UNAME := $(shell uname -s)

	ifeq ($(UNAME),Linux)
		EXT = so
	endif

	ifeq ($(UNAME),Darwin)
		EXT = dylib
	else
		$(error Unsupported system)
	endif
endif


build:
	go build $(LDFLAGS) -o $(ARTIFACT).$(EXT) -buildmode=$(BUILDMODE) .

release: LDFLAGS = -ldflags "-w -s"
release: build

binary: ARTIFACT = fal
binary: BUILDMODE = exe
binary: build
