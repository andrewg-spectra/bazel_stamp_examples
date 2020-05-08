# Bazel Stamp Examples

Bazel has a concept of injecting build metadata into a binary called `stamp`ing. Using 
`--workspace_status_command program` with `--stamp`, a set of key=value pairs can be communicated
to the build system and used in rules.

It is up to the language and it's associated rules to come up with a way to actually inject the 
`stamp` into a binary. This repo contains some examples for Go and C/C++, our two primary 
languages. 