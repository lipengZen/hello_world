package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorString{s: text}
}

var ErrNamedType = NewError("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == NewError("EOF") {
		fmt.Println("Error: Named Error type")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("struct Error type")
	}
}
