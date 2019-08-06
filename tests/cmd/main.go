package main

import (
	"sln/tests"
)

func main() {
	var (
		conf = tests.ReadConfig("./tests/cmd/config.yaml")
		gen  = tests.NewGenerator(conf)
	)
	gen.PlayScripts()
}
