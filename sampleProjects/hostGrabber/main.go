package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	// "github.com/felixge/netroute/table"
)

func getHostInfo() (string, string, string, string) {
	// Get the Hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Get the Username

	var username string
	if runtime.GOOS == "windows" {
		username = os.Getenv("USERNAME")
	} else {
		username = os.Getenv("USER")
	}

	// Get the IP Address

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	var ipAddress string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ipAddress = ipnet.IP.String()
		}
	}
	if ipAddress == "" {
		panic("Unable to find IP Address for this host.")
	}

	// Get the MAC Address
	interfaceList, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var macAddress string
	for _, iface := range interfaceList {
		if iface.HardwareAddr != nil {
			macAddress = iface.HardwareAddr.String()
			break
		}
	}
	if macAddress == "" {
		panic("Unable to find MAC Address for this host.")
	}
	// Get the Gateway Address and Subnet Mask
	// routes, err := table.RouteTable()
	//if err != nil {
	//	panic(err)
	//}
	// defaultRoute, err := routes.Find(net.IPv4(0, 0, 0, 0))
	//if err != nil {
	//	panic(err)
	//}
	//gateway := defaultRoute.Gateway.String()
	//subnetMask := net.IP(defaultRoute.Mask).String()

	return hostname, username, ipAddress, macAddress
}
func main() {
	hostname, username, ipAddress, macAddress := getHostInfo()
	var i = fmt.Println

	i("Hostname: ", hostname)
	i("Username: ", username)
	i("IP Address: ", ipAddress)
	i("MAC Address: ", macAddress)
	// i("Gateway: ", gateway)
	// i("Subnet Mask: ", subnetMask)
}
