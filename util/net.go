package util

import (
	"fmt"
	"net"
)

// ExternalIP 获取ip
func ExternalIP() net.IP {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrList, err := iface.Addrs()
		if err != nil {
			return nil
		}
		for _, addr := range addrList {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip
		}
	}
	return nil
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func LocalIp() string {
	ip := "127.0.0.1"

	addrList, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Printf("获取本地IP发生异常 : %v\n", err)
	}

	for _, addr := range addrList {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()

				break
			}
		}
	}

	return ip
}
