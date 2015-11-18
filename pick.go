package main

import (
	"flag"
	"fmt"
)

const appVersion = "0.1.0"

func main() {
	version := flag.Bool("v", false, "Print version")
	refresh := flag.Bool("r", false, "Refresh ip ranges")
	config := flag.String("c", "config.json", "Config file")
	output := flag.String("t", "ros", "Specifies the output platform")

	flag.Parse()

	if *version {
		fmt.Println(appVersion)
	} else if *refresh {
		fmt.Println("Refresh")
	}

	fmt.Println("Config file:", *config)

	fmt.Println("Platform:", *output)
}
