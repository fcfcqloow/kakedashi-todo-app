package domainerr

import (
	"fmt"
)

type NotFoundError struct {
	OriginErr error
}

func (n *NotFoundError) Error() string {
	if n.OriginErr != nil {
		return fmt.Errorf("not found %w", n.OriginErr).Error()
	}

	return "not found"
}

func (n *NotFoundError) Unwrap() error {
	return n.OriginErr
}

func IsNotFoundError(err error) bool {
	return isErr[*NotFoundError](err)
}
