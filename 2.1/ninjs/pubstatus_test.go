package ninjs_test

import (
	"fmt"
	"testing"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

func TestPubstatus_Validate(t *testing.T) {
	cases := []struct {
		name        string
		actual      ninjs.Pubstatus
		expectedErr error
	}{
		{
			name:        "canceled",
			actual:      ninjs.PubstatusCanceled,
			expectedErr: nil,
		},
		{
			name:        "usable",
			actual:      ninjs.PubstatusUsable,
			expectedErr: nil,
		},
		{
			name:        "withheld",
			actual:      ninjs.PubstatusWithheld,
			expectedErr: nil,
		},
		{
			name:        "invalid",
			actual:      "invalid",
			expectedErr: fmt.Errorf("invalid value (expected one of canceled, usable, withheld): invalid"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.actual.Validate()
			DiffError(t, c.expectedErr, err)
		})
	}
}
