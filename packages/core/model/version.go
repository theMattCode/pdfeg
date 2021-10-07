package model

import "fmt"

type Version struct {
	Major byte
	Minor byte
}

func (v Version) String() string {
	return fmt.Sprintf("%0d.%0d", v.Major, v.Minor)
}
