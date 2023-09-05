package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	// "github.com/felixge/netroute/table"
)

func getVPNInfo() (string, string) {
	cmd := exec.Command("scutil", "--nc", "list")
	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(out)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	rawOutput := ""
	for scanner.Scan() {
		rawOutput += scanner.Text() + "\n"
	}
	cmd.Wait()

	vpnName := ""
	vpnIP := ""

	if strings.Contains(rawOutput, "Connected") {
		lines := strings.Split(rawOutput, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Connected") {
				vpnName = strings.Split(line, " (Connected by ")[0]
			}
			if strings.Contains(line, "ServerAddress : ") {
				vpnIP = strings.Replace(strings.Split(line, "ServerAddress : ")[1], "Address : ", "", -1)
			}
		}
	}
	return vpnName, vpnIP
}
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

	// Get the VPN information

	return hostname, username, ipAddress, macAddress
}
func main() {
	hostname, username, ipAddress, macAddress := getHostInfo()

	vpnName, vpnIP := getVPNInfo()

	// Print Results
	var i = fmt.Println

	i("Hostname: ", hostname)
	i("Username: ", username)
	i("IP Address: ", ipAddress)
	i("MAC Address: ", macAddress)
	if runtime.GOOS == "darwin" {
		i("VPN Name:", vpnName)
		i("VPN IP Address:", vpnIP)
	} else if runtime.GOOS == "windows" {
		i("VPN IP Address:", vpnIP)
	}
	// i("Gateway: ", gateway)
	// i("Subnet Mask: ", subnetMask)
}
