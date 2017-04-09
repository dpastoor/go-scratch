package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ips := getAllIps()
	for _, ip := range ips {
		if ip.Error != nil {
			log.Printf("Error with ip %s: %s", ip.Name, ip.Error)
		}
		fmt.Printf("%s: %s\n", ip.Name, ip.IP)
	}
}

type ipinfo struct {
	IP    net.IP
	Name  string
	Error error
}

func getAllIps() []ipinfo {
	ifaces, _ := net.Interfaces()
	var ips []ipinfo
	// handle err
	for _, i := range ifaces {
		var ip net.IP
		addrs, err := i.Addrs()
		if err != nil {
			ips = append(ips, ipinfo{Name: i.Name, Error: err})
			continue
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if ipv4 := v.IP.To4(); ipv4 != nil {
					ip = v.IP
				}
			}
		}
		ips = append(ips, ipinfo{IP: ip, Name: i.Name, Error: err})
	}
	return ips
}
