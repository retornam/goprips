package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	prips "github.com/retornam/goprips/prips"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	var cidr string
	flag.StringVar(&cidr, "cidr", "", "CIDR range to expand example 192.168.0.0/24")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if cidr != "" {
		items, err := prips.Hosts(cidr)
		check(err)
		var b bytes.Buffer
		for _, item := range items {
			fmt.Fprintf(&b, "%v\n", item)
		}
		fmt.Print(b.String())
		os.Exit(0)
	} else {
		flag.Usage()
	}
}
