HOST_PLATFORM=$(shell uname | tr '[:upper:]' '[:lower:]')
Q=@

all:
	${Q}bazel build //neopixel:neopixel

upload: 
	${Q}bazel run //neopixel:neopixel

test:

clean:
	${Q}bazel clean --expunge
	
.PHONY: all upload test clean