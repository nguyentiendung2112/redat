load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "redat_main",
    srcs = ["main.go"],
    importpath = "redat",  # matches module path in go.mod
    deps = ["//:core_lib"],
)

go_library(
    name = "core_lib",
    srcs = glob(
        ["core/*go"],
    ),
    importpath = "redat/core",
)

go_binary(
    name = "redat_cli_main",
    srcs = ["cli/cli_main.go"],
    importpath = "redat/cli",
    deps = ["//:core_lib"],
)
