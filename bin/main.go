package main

import "github.com/atotto/clipboard/cmd"

var version = "0.1.5"

func main() {
	cmd.Version = version
	cmd.Execute()
}
