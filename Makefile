HOST_PLATFORM=$(shell uname | tr '[:upper:]' '[:lower:]')
Q=@

all:
	${Q}bazel build //...

test:

clean:
	${Q}bazel clean --expunge

neopixel: 
	${Q}bazel run //neopixel:neopixel

ios:
	${Q}bazel build //ios-led-controller:led-controller
	${Q}ideviceinstaller --uninstall com.robstein.LEDController
	${Q}ideviceinstaller --install bazel-bin/ios-led-controller/led-controller.ipa

.PHONY: all test clean neopixel ios