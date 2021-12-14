package custom_errors

import "fmt"

const (
	Internal     = "internal"
	InvalidInput = "invalid_input"
	Panic        = "panic"

	defaultType = Internal
)

type customError struct {
	errType string
	msg     string
	cause   error
}

func (e customError) Error() string {
	return fmt.Sprintf("%s, type: %s, cause: <%v>", e.msg, e.errType, e.cause)
}

func New(message string, cause error) error {
	return customError{msg: message, errType: Type(cause), cause: cause}
}

func NewWithType(message string, t string, cause error) error {
	return customError{msg: message, errType: t, cause: cause}
}

func Message(e error) string {
	if e == nil {
		return ""
	}
	if v, ok := e.(customError); ok {
		return v.msg
	}
	return e.Error()
}

func Type(e error) string {
	if e == nil {
		return defaultType
	}
	if v, ok := e.(customError); ok {
		return v.errType
	}
	return defaultType
}

func Cause(e error) error {
	if e == nil {
		return nil
	}
	if v, ok := e.(customError); ok {
		return v.cause
	}
	return nil
}
