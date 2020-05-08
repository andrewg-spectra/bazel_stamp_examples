# Go `x_defs`

This example uses the `x_defs` attribute of `go_binary` to inject a value into a variable that has
been initially set to "redacted".

## With `bazel`:

`bazelisk run --stamp --workspace_status_command=$PWD/../version.sh //cmd/go_x_defs`

## With `go`:

`export $(../version.sh | tr ' ' =)`
`go run -ldflags "-X github.com/andrewg-spectra/bazel_stamp_examples/go_x_defs/internal/version.BuildVersion=$BUILD_VERSION" cmd/go_x_defs/main.go`
