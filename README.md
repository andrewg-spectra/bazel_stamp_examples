# Bazel Stamp Examples

Bazel has a concept of injecting build metadata into a binary called stamping. Using 
`--workspace_status_command program` with `--stamp`, a set of key=value pairs can be communicated
to the build system and used in rules.

It is up to the language and it's associated rules to come up with a way to actually inject the 
stamp into a binary. This repo contains some examples for Go and C/C++.

## Concepts

There are two different styles of stamps:

- *stable*: if the value of the stamp changes, re-link the binary.
- *volatile*: if the the value of the stamp changes, do not re-link the binary.

To use a *stable* stamp, prefix the name of the key with `STABLE_`. Otherwise, the stamp will be
considered volatile. For the purposes of setting the build version, a volatile stamp is desired.

## Strategies

Regardless of the language, there are two general strategies to stamp binaries:

1. Link-time definitions
1. Code generation

Since code generation is a heavy handed strategy that includes developing and maintaining code 
generators, prefer using link-time definitions when supported.
