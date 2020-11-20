package types

// DriveConfig Drive Config Type
type DriveConfig struct {
	Drives []string
}

// DiskStatus Type for Drive Info
type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}