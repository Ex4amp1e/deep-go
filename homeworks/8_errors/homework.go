package homework8

import (
	"fmt"
	"strings"
)

type MultiError struct {
	list []error
}

func (e *MultiError) Error() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d errors occurred:\n", len(e.list))
	for _, value := range e.list {
		fmt.Fprintf(&sb, "\t* %v", value)
	}
	sb.WriteByte('\n')
	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	errMulti := &MultiError{make([]error, 0)}

	me, ok := err.(*MultiError)
	if !ok {
		errMulti.list = append(errMulti.list, errs...)
	} else {
		errMulti.list = append(me.list, errs...)
	}

	return errMulti
}
