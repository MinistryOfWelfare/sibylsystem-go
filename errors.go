package SibylSystem

import (
	"errors"
	"fmt"
)

var ErrInvalidToken = errors.New("your token seems to be invalid")

func (e *EndpointError) Error() string {
	return fmt.Sprintf("An error encountered while executing a method:\n Error Code: %d\n Message: %s\n Origin: %s\n Date: %s", e.ErrorCode, e.Message, e.Origin, e.Date)
}
