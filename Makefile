HOST_PLATFORM=$(shell uname | tr '[:upper:]' '[:lower:]')
Q=@

all:
	${Q}bazel build //neopixel:neopixel

upload: 
	${Q}bazel run //neopixel:neopixel

test:

clean:
	${Q}bazel clean --expunge

ios:
	${Q}bazel build --host_force_python=PY2 //ios-led-controller:led-controller
	${Q}ideviceinstaller --install bazel-bin/ios-led-controller/led-controller.ipa
	
.PHONY: all upload test clean ios