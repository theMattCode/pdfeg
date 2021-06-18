package objects

import (
	"encoding/hex"
	"fmt"
)

type String interface {
	IsLiteral() bool
	IsHexadecimal() bool
	String() string
}

type LiteralString struct {
	*Reference
	Value string
}

func (s LiteralString) String() string {
	return s.Value
}

func (s LiteralString) IsLiteral() bool {
	return true
}

func (s LiteralString) IsHexadecimal() bool {
	return false
}

func (s LiteralString) AsASCIIBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("(%s)", s.Value)), nil
}

type HexadecimalString struct {
	*Reference
	Value string
}

func (s HexadecimalString) String() string {
	return s.Value
}

func (s HexadecimalString) IsLiteral() bool {
	return false
}

func (s HexadecimalString) IsHexadecimal() bool {
	return true
}

func (s HexadecimalString) AsASCIIBytes() ([]byte, error) {
	hexString := hex.EncodeToString([]byte(s.Value))
	return []byte(fmt.Sprintf("<%s>", hexString)), nil
}
