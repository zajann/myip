package myip

import "testing"

func TestGetIPv4(t *testing.T) {
	ipv4 := GetIPv4("lo")
	t.Log(ipv4)
}

func TestGetIPv6(t *testing.T) {
	ipv6 := GetIPv6("lo")
	t.Log(ipv6)
}
