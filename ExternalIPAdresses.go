package main

import (
	"errors"
	"net"
)

// ExternalIPAdresses ...
func ExternalIPAdresses() ([]string, error) {
	var ipAddresses []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, interf := range interfaces {

		if interf.Flags&net.FlagUp == 0 {
			continue
		}

		if interf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := interf.Addrs()
		if err != nil {
			return nil, err
		}
		for _, address := range addresses {

			var ip net.IP

			switch a := address.(type) {
			case *net.IPNet:
				ip = a.IP
			case *net.IPAddr:
				ip = a.IP
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			ipAddresses = append(ipAddresses, ip.String())
		}
	}
	if len(ipAddresses) == 0 {
		return ipAddresses, errors.New("No interfaces up or connected")
	}
	return ipAddresses, nil
}
