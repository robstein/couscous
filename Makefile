HOST_PLATFORM=$(shell uname | tr '[:upper:]' '[:lower:]')
Q=@

all:
	${Q}bazel build $(shell bazel query 'filter(".*", kind(".*", //...)) except attr("tags", "manual", //...)')

test:

clean:
	${Q}bazel clean --expunge

neopixel: 
	${Q}bazel run //neopixel-driver:neopixel-driver

server: 
	${Q}bazel run //server:server

ios:
	${Q}bazel build //ios-led-controller:led-controller
	${Q}ideviceinstaller --uninstall com.robstein.LEDController
	${Q}ideviceinstaller --install bazel-bin/ios-led-controller/led-controller.ipa

.PHONY: all test clean neopixel server ios