package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var network = flag.String(
	"network",
	"tcp",
	"network type to dial with (e.g. unix, tcp)",
)

var addr = flag.String(
	"addr",
	":8080",
	"address to dial",
)

var timeout = flag.Duration(
	"timeout",
	1*time.Second,
	"dial timeout",
)

func main() {
	flag.Parse()

	conn, err := net.DialTimeout(*network, *addr, *timeout)
	if err != nil {
		fmt.Println("healthcheck failed")
		os.Exit(1)
	}

	conn.Close()

	fmt.Println("healthcheck passed")
	os.Exit(0)
}
