workspace(name = "couscous")

load("//:version_check.bzl", "check_bazel_version_at_least")
check_bazel_version_at_least("0.27.1")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository", "new_git_repository")

skylib_version = "0.8.0"
http_archive(
    name = "bazel_skylib",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib.{}.tar.gz".format (skylib_version, skylib_version),
    sha256 = "2ef429f5d7ce7111263289644d233707dba35e39696377ebab8b0bc701f7818e",
)

apple_support_version = "0.6.0"
http_archive(
    name = "build_bazel_apple_support",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/apple_support/releases/download/{}/apple_support.{}.tar.gz".format (apple_support_version, apple_support_version),
    sha256 = "7356dbd44dea71570a929d1d4731e870622151a5f27164d966dda97305f33471",
)

load("@build_bazel_apple_support//lib:repositories.bzl", "apple_support_dependencies")

apple_support_dependencies()

rules_apple_version = "0.17.2"
http_archive(
    name = "build_bazel_rules_apple",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/rules_apple/releases/download/{}/rules_apple.{}.tar.gz".format (rules_apple_version, rules_apple_version),
    sha256 = "6efdde60c91724a2be7f89b0c0a64f01138a45e63ba5add2dca2645d981d23a1",
)

rules_go_version = "0.18.6"
http_archive(
    name = "io_bazel_rules_go",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/rules_go/releases/download/{}/rules_go-{}.tar.gz".format (rules_go_version, rules_go_version),
    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d",
)

git_repository(
    name = "bazel_gazelle",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
    commit = "e443c54b396a236e0d3823f46c6a931e1c9939f2", 
    shallow_since = "1551292640 -0800",
)

git_repository(
    name = "bazel_buildtools",
    remote = "https://github.com/bazelbuild/buildtools.git",
    commit = "eb1a85ca787f0f5f94ba66f41ee66fdfd4c49b70", 
    shallow_since = "1559821206 +0200",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_register_toolchains(nogo = "@//:my_nogo")

go_rules_dependencies()

load("@bazel_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

git_repository(
    name = "platformio_rules",
    remote = "http://github.com/mum4k/platformio_rules.git",
    commit = "ee2c5a9e302520270ae5455a572357e5b30a7d33",
    shallow_since = "1565179280 -0400",
)

new_git_repository(
    name = "com_github_FastLED_FastLED",
    remote = "https://github.com/FastLED/FastLED.git",
    commit = "a346de18a09ad4471c6cc8bbefe4eb8bcf863d32",
    shallow_since = "1562172155 -0700",
    build_file = "com_github_FastLED_FastLED.BUILD",
)
