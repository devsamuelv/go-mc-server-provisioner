package drives

import (
	"fmt"
	"os/exec"
	"strconv"
)

// CreatePartition create's a partition then returns a location
// Make the partition name the driveId
func CreatePartition(disk string, partitionName string, size int) (location string) {
	fileSysytemType := "ext4"
	driveType := "secondary"

	createCmd := exec.Command("bin/sh", "-c", "sudo parted "+disk+"; mklabel msdos; mkpart "+driveType+" "+fileSysytemType+" 0 "+strconv.Itoa(size)+"GB ")

	err := createCmd.Run()

	if err != nil {
		fmt.Print(err.Error())
	}

	return ""
}
