/**
* @file goerror.go
* @brief error with errno and msg
* @author ligang
* @date 2020-03-17
 */

package goerror

import (
	"strconv"
)

type Error struct {
	errno int
	msg   string
}

func New(errno int, msg string) *Error {
	return &Error{
		errno: errno,
		msg:   msg,
	}
}

func (e *Error) Error() string {
	result := "errno: " + strconv.Itoa(e.errno) + ", "
	result += "msg: " + e.msg

	return result
}

func (e *Error) Errno() int {
	return e.errno
}

func (e *Error) Msg() string {
	return e.msg
}
