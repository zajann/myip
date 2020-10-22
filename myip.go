package myip

import (
	"net"
)

// GetIPv4 returns the IPv4 address by specified by interface name.
// if the interface name is incorrect or there is no ip, return empty string
func GetIPv4(name string) string {
	var ipv4 net.IP
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return ""
	}
	addrs, err := iface.Addrs()
	if err != nil || len(addrs) == 0 {
		return ""
	}
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			ipv4 = v.IP
		}
	}
	if ipv4.To4() == nil {
		return ""
	}
	return ipv4.String()
}

// GetIPv6 returns the IPv6 address by specified by interface name.
// if the interface name is incorrect or there is no ip, return empty string
func GetIPv6(name string) string {
	var ipv6 net.IP
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return ""
	}
	addrs, err := iface.Addrs()
	if err != nil || len(addrs) == 0 {
		return ""
	}
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPAddr:
			ipv6 = v.IP
		}
	}
	if ipv6.To4() != nil {
		return ""
	}
	return ipv6.String()
}
