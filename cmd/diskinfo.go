package cmd

import (
	"fmt"
	"log"
	"runtime"

	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

var diskInfoCmd = &cobra.Command{
	Use:   "diskinfo",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		GetDiskInfo()
	},
}

func init() {
	rootCmd.AddCommand(diskInfoCmd)
}

func GetDiskInfo(drive ...string) error {
	// Retrieves hard disk information

	targetDrive := "C:" // Default for windows
	if runtime.GOOS != "windows" {
		targetDrive = "/" // Default for other systems (Linux, macOS, etc.)
	}

	// We just want to perform operations on a single drive
	if len(drive) > 0 {
		targetDrive = drive[0]
	}

	usage, err := disk.Usage(targetDrive)
	if err != nil {
		log.Fatal(err)
	}

	totalDiskUsage := usage.Total / 1024 / 1024 / 1024
	freeDiskSpace := usage.Free / 1024 / 1024 / 1024
	usedDiskSpace := usage.Used / 1024 / 1024 / 1024
	usedDiskSpacePercentage := usage.UsedPercent
	filesystemType := usage.Fstype
	path := usage.Path

	fmt.Println("Disk Information:")
	fmt.Printf("  Total: %v GB\n", totalDiskUsage)
	fmt.Printf("  Free: %v GB\n", freeDiskSpace)
	fmt.Printf("  Used: %v GB (%.2f%%)\n", usedDiskSpace, usedDiskSpacePercentage)
	fmt.Printf("  Filesystem: %s\n", filesystemType)
	fmt.Printf("  Mount Point: %s\n", path)
	return nil
}
