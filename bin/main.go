package main

import "github.com/kawana77b/go-clip/cmd"

var version = "0.1.5"

func main() {
	cmd.Version = version
	cmd.Execute()
}
