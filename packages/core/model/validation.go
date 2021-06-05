package model

import (
	"fmt"
	"strconv"
)

const (
	PositiveIntegerMessage = "integer must greater than or equal to 1"
	NonNegativeIntegerMessage = "integer must greater than or equal to 0"
)

type ValidationError struct {
	Context string
	Message string
	Value   string
}

func ValidatePositiveInteger(context string, i int) *ValidationError {
	if i <= 0 {
		return &ValidationError{context, PositiveIntegerMessage, strconv.Itoa(i)}
	}
	return nil
}

func ValidateNonNegativeInteger(context string, i int) *ValidationError {
	if i < 0 {
		return &ValidationError{context, NonNegativeIntegerMessage, strconv.Itoa(i)}
	}
	return nil
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s: %s", err.Context, err.Message, err.Value)
}
