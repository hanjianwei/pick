package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hanjianwei/pick/netutils"
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

	ip1 := net.IPv4(192, 168, 0, 0)
	ip2 := net.IPv4(192, 168, 0, 255)
	d := netutils.IPDistance(ip2, ip1)
	f := netutils.IPDistance(ip1, ip2)

	fmt.Println(d)
	fmt.Println(f)

	//
	// ipr.addIPNet(ip1)
	// ipr.addIPNet(ip2)
}
