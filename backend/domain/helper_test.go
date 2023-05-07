package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testError(t *testing.T, want, err error) {
	t.Helper()
	if err != nil && want != nil {
		assert.EqualError(t, err, want.Error(), fmt.Sprintf("expected %s, got %s", want.Error(), err.Error()))
		return
	}

	assert.NoError(t, err)
}
