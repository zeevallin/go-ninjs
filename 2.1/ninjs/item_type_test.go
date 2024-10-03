package ninjs_test

import (
	"fmt"
	"testing"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

func TestItemType_Validate(t *testing.T) {
	cases := []struct {
		name        string
		actual      ninjs.ItemType
		expectedErr error
	}{
		{
			name:        "audio",
			actual:      ninjs.ItemTypeAudio,
			expectedErr: nil,
		},
		{
			name:        "component",
			actual:      ninjs.ItemTypeComponent,
			expectedErr: nil,
		},
		{
			name:        "composite",
			actual:      ninjs.ItemTypeComposite,
			expectedErr: nil,
		},
		{
			name:        "graphic",
			actual:      ninjs.ItemTypeGraphic,
			expectedErr: nil,
		},
		{
			name:        "picture",
			actual:      ninjs.ItemTypePicture,
			expectedErr: nil,
		},
		{
			name:        "text",
			actual:      ninjs.ItemTypeText,
			expectedErr: nil,
		},
		{
			name:        "video",
			actual:      ninjs.ItemTypeVideo,
			expectedErr: nil,
		},
		{
			name:        "invalid",
			actual:      "invalid",
			expectedErr: fmt.Errorf("invalid value (expected one of audio, component, composite, graphic, picture, text, video): invalid"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.actual.Validate()
			DiffError(t, c.expectedErr, err)
		})
	}
}
