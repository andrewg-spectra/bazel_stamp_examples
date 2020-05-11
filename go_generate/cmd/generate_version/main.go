package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//go:generate go run . -in=../../internal/version/version.go -out=../../internal/version/version.gen.go -ver=$BUILD_VERSION

var (
	in  = flag.String("in", "", "input")
	out = flag.String("out", "", "output")
	ver = flag.String("ver", "", "version")
)

func main() {
	flag.Parse()

	if len(*ver) == 0 {
		log.Fatal("-ver flag required")
	}

	input, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = append(lines, "// BuildVersion of binary")
	lines = append(lines, fmt.Sprintf("var BuildVersion = \"%s\"", *ver))

	output := strings.Join(lines, "\n")
	if err := ioutil.WriteFile(*out, []byte(output), 0644); err != nil {
		log.Fatal(err)
	}
}
