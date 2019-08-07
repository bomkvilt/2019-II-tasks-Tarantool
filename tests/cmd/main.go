package main

import (
	"os"
	"sln/tests"
)

func main() {
	var (
		conf = tests.ReadConfig("./tests/cmd/config.yaml")
		gen  = tests.NewGenerator(conf)
	)
	if !gen.PlayScripts() {
		os.Exit(1)
	}
}
