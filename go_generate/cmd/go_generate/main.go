package main

import (
	"fmt"

	"github.com/andrewg-spectra/bazel_stamp_examples/go_generate/internal/version"
)

func main() {
	fmt.Printf("v1.2.3.%s\n", version.BuildVersion)
}
