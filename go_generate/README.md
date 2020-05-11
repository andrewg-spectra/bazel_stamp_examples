# Go generate

This examples uses code generation to write Go source code containing a build number. The generated
source code is compiled into the binary. 

## With `bazel`:


## With `go`:

`export $(../version.sh | tr ' ' =)`

`go generate cmd/generate_version/main.go`

`go run cmd/go_generate/main.go`

## With `bazel`

`bazelisk run --stamp --workspace_status_command=$PWD/../version.sh //cmd/go_generate`
