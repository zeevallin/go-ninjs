package ninjs_test

import (
	"fmt"
	"testing"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

func TestRepresentationType_Validate(t *testing.T) {
	cases := []struct {
		name        string
		actual      ninjs.RepresentationType
		expectedErr error
	}{
		{
			name:        "full",
			actual:      ninjs.RepresentationTypeFull,
			expectedErr: nil,
		},
		{
			name:        "partial",
			actual:      ninjs.RepresentationTypePartial,
			expectedErr: nil,
		},
		{
			name:        "invalid",
			actual:      "invalid",
			expectedErr: fmt.Errorf("invalid value (expected one of full, partial): invalid"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.actual.Validate()
			DiffError(t, c.expectedErr, err)
		})
	}
}
