# Go `x_defs`

This example uses the `x_defs` attribute of `go_binary` to inject a value into a variable that has
been initially set to "redacted".

With `bazel`:

`bazelisk run --stamp --workspace_status_command=$PWD/../version.sh //cmd`

With `go`:

`go build