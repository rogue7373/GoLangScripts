package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	var ip net.IP
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				break
			}
		}
	}

	iface, err := net.InterfaceByIndex(1) // Using the first interface by default
	if err != nil {
		fmt.Println(err)
		return
	}

	addrs, err = iface.Addrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	var subnet net.IPMask
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				subnet = ipnet.Mask
				break
			}
		}
	}
	// The below code that is commented out is not working to pull the gateway and MAC address for the local host. Need to investigate a better way to pull these details.

	// defaultRoute, err := netlink.RouteGet(net.IPv4(0, 0, 0, 0))
	// if err != nil {
	//	fmt.Println(err)
	//	return
	// }

	// gatewayIP := defaultRoute[0].Gw

	// haddr, err := iface.HardwareAddr()
	// if err != nil {
	//	fmt.Println(err)
	//	return
	// }

	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("IP Address: %s\n", ip.String())
	fmt.Printf("Subnet Mask: %s\n", subnet.String())
	// fmt.Printf("Gateway Address: %s\n", gatewayIP.String())
	// fmt.Printf("MAC Address: %s\n", haddr.String())
}
