package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip, err := getIP()
	if err != nil {
		log.Printf("error getting IP: %s", err)
	}
	fmt.Println(ip)
}
func getIP() (net.IP, error) {
	ifaces, _ := net.Interfaces()
	var ip net.IP
	// handle err
	for _, i := range ifaces {
		if i.Name == "eth0" {
			addrs, err := i.Addrs()
			if err != nil {
				return ip, err
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
		}
	}
	return ip, nil
}
