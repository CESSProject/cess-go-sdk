package utils

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/debug"
)

// RecoverError is used to record the stack information of panic
func RecoverError(err interface{}) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", "[panic]")
	fmt.Fprintf(buf, "%v\n", err)
	if debug.Stack() != nil {
		fmt.Fprintf(buf, "%v\n", string(debug.Stack()))
	}
	return errors.New(buf.String())
}
