package main

import cobracli "github.com/felixfwu/datatop/cmd/cobracli/cmd"

var version string

func main() {
	cobracli.Execute(version)
}
