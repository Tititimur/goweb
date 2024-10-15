package main

import (
	"fmt"

	"github.com/tititimur/goweb/server"
)

var VersionCode = ""

func main() {
	fmt.Printf("### Start GoWeb\n")
	server.Start()
}
