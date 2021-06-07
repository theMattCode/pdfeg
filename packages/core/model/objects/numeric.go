package objects

import (
	"fmt"
	"strconv"
)

// Integer represents PDF's integer object. See ISO 32000-2:2017, 7.3.3
type Integer struct {
	*Reference
	Value int64
}

func (i Integer) AsBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", i.Value)), nil
}

// Real represents PDF's real object. See ISO 32000-2:2017, 7.3.3
type Real struct {
	*Reference
	Value float64
}

func (r Real) AsBytes() ([]byte, error) {
	return []byte(strconv.FormatFloat(r.Value, 'f', -1, 64)), nil
}

type Number struct {
	*Integer
	*Real
}
