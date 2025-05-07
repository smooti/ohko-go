package main

import (
	"fmt"

	"github.com/yusufpapurcu/wmi"
)

// Define a struct that matches the Win32_OperatingSystem WMI class
// NOTE: The field name must match the WMI property name
type Win32_OperatingSystem struct {
	SerialNumber string
}

func main() {

	fmt.Println(getWindowsSerialNumber())

}

func getWindowsSerialNumber() string {
	// Return the serial number using WMI
	var (
		biosInfo []Win32_OperatingSystem
		query    = "SELECT SerialNumber FROM Win32_OperatingSystem"
	)

	// Query BIOS information
	err := wmi.Query(query, &biosInfo)
	if err != nil {
		fmt.Printf("WMI Query failed: %v", err)
	}

	// Check if biosInfo was retrieved
	if len(biosInfo) > 0 {
		return biosInfo[0].SerialNumber
	} else {
		return "No BIOS info found."
	}
}
