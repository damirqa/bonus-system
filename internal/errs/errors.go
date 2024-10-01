package errs

import "fmt"

type ErrUserAlreadyExists struct {
	Err error
}

func NewErrUserAlreadyExists(err error) *ErrUserAlreadyExists {
	return &ErrUserAlreadyExists{Err: err}
}

func (e *ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("Unique constraint error: %v", e.Err)
}
