package errordemo

import "fmt"

type CustomError struct {
}

func (e *CustomError) Error() string {
	return "this is a custom error"
}

// If you're using custom error types, don't return underlying type, always return error type.,
func returnsCustomError(pleaseFail bool) *CustomError {
	if pleaseFail {
		return &CustomError{}
	}
	return nil
}

// Unexpectedly it looks like returnsCustomError() returned an error even though it returned nil.
// This is a variant of nil interface pitfall.
// #note link to that chapter
// Uninitialized err value is nil because it's the zero value of interface types.
// returnsCustomError() returns nil value of CustomError struct.
// This nil value is assigned to err variable. This is equivalent to:
// var err error
// err = (*CustomError)(nil)
// At this point err is no longer equal to zero value of interface type. It's value is *CustomError whose value is nil.
func caller() {
	var err error
	if err == nil {
		fmt.Printf("err is nil, as expected\n")
	}

	err = returnsCustomError(false)
	if err != nil {
		fmt.Printf("returnCustomError() failed with '%s'\n", err)
	}
}

func Run() {
	caller()
}
