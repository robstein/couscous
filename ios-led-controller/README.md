## How to build

Prerequisites:
 - Update `.bazelrc` with your iOS signing cert name
 - Update `ios-led-controller/profile.mobileprovision` with your mobileprovision
 - Update the bundle id in `ios-led-controller/Info.plist` and `ios-led-controller/BUILD.bazel`

Build:
 - `make ios`
