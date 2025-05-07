package info

import (
	"fmt"
	"log"

	"github.com/elastic/go-sysinfo"
	"github.com/spf13/cobra"
)

var osInfoCmd = &cobra.Command{
	Use:   "osinfo",
	Short: "Get OS information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		GetOSInfo()
	},
}

func init() {
	InfoCmd.AddCommand(osInfoCmd)
}

func GetOSInfo() error {
	// Retrieves operating system information.

	// Return error if unable to get host information
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatal(err)
	}

	info := host.Info()
	hostname := info.Hostname
	osName := info.OS.Name
	osFamily := info.OS.Family
	osVersion := info.OS.Version
	osBuild := info.OS.Build
	architecture := info.Architecture

	fmt.Println("OS Information:")
	fmt.Printf("  Hostname: %s\n", hostname)
	fmt.Printf("  OS Family: %s\n", osFamily)
	fmt.Printf("  OS Name: %s\n", osName)
	fmt.Printf("  OS Version: %s\n", osVersion)
	fmt.Printf("  OS Build: %s\n", osBuild)
	fmt.Printf("  Architecture: %s\n", architecture)
	return nil
}
