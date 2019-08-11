# couscous
an LED strip controller

## How to build

Prerequisites:
 - Buy an [ESP8266](https://www.amazon.com/gp/product/B010N1SPRK) board, a [NeoPixel LED strip](https://www.amazon.com/gp/product/B077F8SQBV), and some [jumper wires](https://www.amazon.com/gp/product/B01LZF1ZSZ).
 - Install the [CP2102 driver](https://www.silabs.com/products/development-tools/software/usb-to-uart-bridge-vcp-drivers) 
 - Install [platformio](http://docs.platformio.org/en/latest/installation.html#install-shell-commands) which is a cross-platform IoT ecosystem that enables you to develop for microcontrollers without having to use the Arduino IDE
 - Install the [bazel](https://github.com/bazelbuild/bazel/releases/tag/0.27.2) build tool

Build:
 - `make`

Upload to your board:
 - `make upload`
