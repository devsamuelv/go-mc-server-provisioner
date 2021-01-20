package types

// Config sys config type
type Config struct {
	Drives []string `json:"drives"`
	DB     struct {
		URL string `json:"url"`
		Key string `json:"Key"`
	}
	MountedLocation string `json:"mounted_location"`
	Server          struct {
		Port string
	}
	Partitions []Partition
	Cert       struct {
		Public  string `json:"public"`
		Private string `json:"private"`
	}
}

// Partition Type For Partitions In
type Partition struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	UID      string `json:"uid"`
	DriveID  string `json:"drive_id"`
}

// DiskStatus Type for Drive Info
type DiskStatus struct {
	All uint64
	// Used uint64
	Free uint64
}
