# couscous
an LED strip controller

## How to build

 - Buy an [ESP8266](https://www.amazon.com/gp/product/B010N1SPRK) board
 - Install the [CP2102](https://www.silabs.com/products/development-tools/software/usb-to-uart-bridge-vcp-drivers) driver
 - Install [platformio](http://docs.platformio.org/en/latest/installation.html#install-shell-commands) which is a cross-platform IoT ecosystem that enables you to develop for microcontrollers without having to use the Arduino IDE
 - Install [bazel](https://github.com/bazelbuild/bazel/releases/tag/0.27.2)
 - Run `make` to build
 - Run `make upload` to upload to your board
