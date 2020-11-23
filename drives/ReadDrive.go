package drives

import (
	"main/types"
	"syscall"
)

const (
	// B Byte
	B = 0.0009765625

	// KB kilobyte
	KB = 1024 / B

	// MB MegaByte
	MB = 1024 / KB

	// GB GigaByte
	GB = 1024 / MB
)

// ReadSpace Read Drive Space
func ReadSpace(name string) (info types.DiskStatus) {
	var stat syscall.Statfs_t

	wd := name

	err := syscall.Statfs(wd, &stat)

	if err != nil {
		panic("Drive " + name + " Doesn't Exist!")
	}

	info.Free = stat.Bavail * uint64(stat.Bsize)
	info.All = stat.Blocks * uint64(stat.Bsize)

	return info
}
