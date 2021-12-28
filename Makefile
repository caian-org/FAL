.DEFAULT_GOAL := build


init:
	@rm -rf build
	@mkdir build
	@cd build && conan install .. --build=missing

.PHONY: build
build:
	@cd build && cmake .. && cmake --build . --config Release

.PHONY: tests
tests:
	@printf "\n%s\n" "--- python package tests ---"
	@python3 tests/python/main.py
	@printf "\n%s\n" "--- javascript package tests ---"
	@node tests/javascript/index.js
