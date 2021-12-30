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
	endif
endif

LIBFILENAME = libfal.$(EXT)

build:
	go build $(LDFLAGS) -o $(LIBFILENAME) -buildmode=$(BUILDMODE) .

release: LDFLAGS = -ldflags "-w -s"
release: build
	@cp $(LIBFILENAME) $(ARTIFACT)

binary: ARTIFACT = fal
binary: BUILDMODE = exe
binary: build
