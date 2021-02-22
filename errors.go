package vctr

import "fmt"

type ResponseError struct {
	Code    int
	Message string
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("[%d] %s", r.Code, r.Message)
}
