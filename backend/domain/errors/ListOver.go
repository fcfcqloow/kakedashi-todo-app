package domainerr

import (
	"fmt"
)

type ListOverError struct {
	OriginErr error
}

func (n *ListOverError) Error() string {
	if n.OriginErr != nil {
		return fmt.Errorf("list over %w", n.OriginErr).Error()
	}

	return "list over"
}

func (n *ListOverError) Unwrap() error {
	return n.OriginErr
}

func IsListOverError(err error) bool {
	return isErr[*ListOverError](err)
}
