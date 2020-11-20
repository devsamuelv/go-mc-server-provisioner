package util

import (
	"fmt"
	"math"
)

func ByteCountSI(b int64) float64 {
    const unit = 1000
    if b < unit {
        return 0;
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
	}

	if exp == 2 {
		return math.Floor(float64(b)/float64(div));
	}

    return 0;
}

func ByteToGigaByte(bytes float64) (gb int) {
	// KB kilobyte
	KB := 1024 / bytes

	// MB MegaByte
	MB := 1024 / KB

	// GB GigaByte
	GB := 1024 / MB

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	return int(bytes / float64(GB))
}