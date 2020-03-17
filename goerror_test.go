package goerror

import (
	"fmt"
	"testing"
)

func TestGoerror(t *testing.T) {
	e := New(101, "test goerror")

	fmt.Println(e.Errno(), e.Msg())
	fmt.Println(e.Error())
}
