package rprips

import (
	"fmt"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Hosts(cidr string) ([]string, error) {
	ipitem, ipcidr, err := net.ParseCIDR(cidr)
	check(err)
	var ips []string
	for ip := ipitem.Mask(ipcidr.Mask);  ipcidr.Contains(ip); increase(ip) {
		ips = append(ips,ip.String())
	}
	return ips, nil
}

func increase(ip net.IP) {
	for i := len(ip) -1; i >=0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}
