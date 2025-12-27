load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "redat_main",
    srcs = ["main.go"],
    deps = ["//:core_lib"],
)

go_library(
    name = "core_lib",
    srcs = glob(
        ["core/*go"],
    ),
)
