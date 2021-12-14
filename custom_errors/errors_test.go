package custom_errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cause := fmt.Errorf("some err")
	err := New("some message", cause)

	v, _ := err.(customError)

	a := assert.New(t)
	a.Equal(customError{msg: "some message", errType: defaultType, cause: cause}, v)
}

func TestNewWithType(t *testing.T) {
	cause := New("some message", nil)
	err := NewWithType("some message", "some_type", cause)

	v, _ := err.(customError)

	a := assert.New(t)
	a.Equal(customError{msg: "some message", errType: "some_type", cause: cause}, v)
}

func TestCustomError_Error(t *testing.T) {
	cause := fmt.Errorf("some cause")
	err := NewWithType("some message", "some_type", cause)

	s := err.Error()

	a := assert.New(t)
	a.Equal("some message, type: some_type, cause: <some cause>", s)
}
