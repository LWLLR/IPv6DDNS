package main

import (
	"fmt"
	"net"
)

func GetIPV6() (string, error) {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, addr := range address {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() == nil && ipNet.IP.IsGlobalUnicast() {
			// 判断是否为IPv6地址，且不是回环地址
			return ipNet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("no ipv6 address found")
}
