package types

// Config sys config type
type Config struct {
	Drives []string `json:"drives"`
	Db     struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	Server struct {
		Port string
	}
	Cert struct {
		Public  string `json:"public"`
		Private string `json:"private"`
	}
}

// DiskStatus Type for Drive Info
type DiskStatus struct {
	All uint64
	// Used uint64
	Free uint64
}
