package vctr

import "fmt"

// ResponseError wraps a HTTP response
// error with error response code and
// an API error message.
//
// ResponseError implements the error
// interface.
type ResponseError struct {
	Code    int
	Message string
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("[%d] %s", r.Code, r.Message)
}
