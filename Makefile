.DEFAULT_GOAL := build


.PHONY: build
build:
	g++ \
		-fPIC \
		-fvisibility=hidden \
		-shared \
		-o build/FAL.so \
		FAL/FAL.cpp
