.DEFAULT_GOAL := build


init:
	rm -rf build && mkdir build && cd build && conan install .. --build=missing

.PHONY: build
build:
	cd build && cmake .. && cmake --build . --config Release
