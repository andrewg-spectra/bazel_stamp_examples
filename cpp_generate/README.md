## C++ generate

This examples uses code generation to write a C++ header file defining the build number. The 
generated source code is compiled into the binary. 

## With `bazel`:

`bazelisk run --stamp --workspace_status_command=$PWD/../version.sh //main:cpp_generate`

## With `g++`:

`export $(../version.sh | tr ' ' =)`

`./lib/gen_version.sh $BUILD_VERSION > lib/version.gen.sh`

`g++ -I . main/cpp-generate.cc lib/version.cc`

`./a.out`
