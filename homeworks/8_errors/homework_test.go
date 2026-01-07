package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

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

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occurred:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
