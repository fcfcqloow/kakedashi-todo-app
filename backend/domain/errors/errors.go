package domainerr

import "errors"

func isErr[T error](err error) bool {
	var nfErr T
	if errors.As(err, &nfErr) {
		return true
	}
	return errors.Is(err, nfErr)
}
