// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"ipe"
)

// These variables are generated by the linker
// please see the makefile for mor information.
var (
	version    = "version"
	buildstamp = "buildstamp"
	githash    = "githash"
)

// Main function, initialize the system
func main() {
	var filename = flag.String("config", "config.yml", "Config file location")
	flag.Parse()

	printBanner()

	ipe.Start(*filename)
}

// Print a beautiful banner
func printBanner() {
	fmt.Println("Websocket started")
}
