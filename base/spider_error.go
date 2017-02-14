package base

import (
	"bytes"
	"fmt"
)

// mySpiderError struct
type mySpiderError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

// NewSpiderError return a SpiderError struct.
func NewSpiderError(errType ErrorType, errMsg string) SpiderError {
	return &mySpiderError{errType: errType, errMsg: errMsg}
}

func (ce *mySpiderError) Type() ErrorType {
	return ce.errType
}

func (ce *mySpiderError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

// genFullErrMsg return full message.
func (ce *mySpiderError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Spider Error:")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)

	ce.fullErrMsg = fmt.Sprintf("%s\n", buffer.String())
}
