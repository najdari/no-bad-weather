package errs

import (
	"fmt"
	"runtime"
)

type StackError struct {
	error
	Stack string
}

func (e *StackError) Error() string {
	return fmt.Sprintf("%v\nStack trace:\n%s", e.error, e.Stack)
}

func StackInfo(err error) *StackError {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return &StackError{
		error: err,
		Stack: string(buf[:n]),
	}
}

func Wrap(msg string, err error) error {
	return StackInfo(fmt.Errorf(msg+": %w", err))
}
