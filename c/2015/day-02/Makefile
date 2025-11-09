include ../lib/Makefile

.PHONY: part1 run-part1 build-part1

build-part1: $(LIB_OBJECTS)
	gcc part1.c $(LIB_OBJECTS) -o build/part1

run-part1: build-part1
	./build/part1

part1: run-part1
