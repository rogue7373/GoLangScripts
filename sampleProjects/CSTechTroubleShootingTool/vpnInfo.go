package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Run the "scutil" command to get VPN information
	cmd := exec.Command("scutil", "--nc", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the output to a string
	outputStr := string(output)

	// Split the output into lines
	lines := strings.Split(outputStr, "\n")

	// Iterate through the lines to find the VPN information
	var vpnName string
	var ipAddress string
	for _, line := range lines {
		if strings.HasPrefix(line, "Connected") {
			// Extract the VPN name from the "Connected" line
			vpnName = strings.TrimSpace(strings.TrimPrefix(line, "Connected"))
		}
		if strings.HasPrefix(line, "Server address:") {
			// Extract the IP address from the "Server address" line
			ipAddress = strings.TrimSpace(strings.TrimPrefix(line, "Server address:"))
		}
	}

	// Print the VPN name and IP address
	fmt.Println("VPN Name:", vpnName)
	fmt.Println("IP Address:", ipAddress)
}
