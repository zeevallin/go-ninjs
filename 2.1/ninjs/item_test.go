package ninjs_test

import (
	"fmt"
	"testing"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

func TestItem_Validate(t *testing.T) {
	cases := []struct {
		name        string
		actual      ninjs.Item
		expectedErr error
	}{
		{
			name:        "when object type is valid",
			actual:      ninjs.Item{Type: ninjs.ItemTypeAudio},
			expectedErr: nil,
		},
		{
			name:        "when object type is invalid",
			actual:      ninjs.Item{Type: "invalid"},
			expectedErr: fmt.Errorf("error in type: invalid value (expected one of audio, component, composite, graphic, picture, text, video): invalid"),
		},
		{
			name:        "when object type is empty",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object representation type is valid",
			actual:      ninjs.Item{RepresentationType: ninjs.RepresentationTypePartial},
			expectedErr: nil,
		},
		{
			name:        "when object representation type is invalid",
			actual:      ninjs.Item{RepresentationType: "invalid"},
			expectedErr: fmt.Errorf("error in representation type: invalid value (expected one of full, partial): invalid"),
		},

		{
			name:        "when object representation type is empty",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object pubstatus is valid",
			actual:      ninjs.Item{Pubstatus: ninjs.PubstatusCanceled},
			expectedErr: nil,
		},
		{
			name:        "when object pubstatus is invalid",
			actual:      ninjs.Item{Pubstatus: "invalid"},
			expectedErr: fmt.Errorf("error in pubstatus: invalid value (expected one of canceled, usable, withheld): invalid"),
		},
		{
			name:        "when object pubstatus is empty",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object has valid bodies",
			actual:      ninjs.Item{Bodies: []ninjs.Body{{Value: "value"}}},
			expectedErr: nil,
		},
		{
			name:        "when object has invalid bodies",
			actual:      ninjs.Item{Bodies: []ninjs.Body{{Value: ""}}},
			expectedErr: fmt.Errorf("error in body 0: field value: required"),
		},
		{
			name:        "when object has both valid and invalid bodies",
			actual:      ninjs.Item{Bodies: []ninjs.Body{{Value: "value"}, {Value: ""}}},
			expectedErr: fmt.Errorf("error in body 1: field value: required"),
		},
		{
			name:        "when object has empty bodies",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object has valid renditions",
			actual:      ninjs.Item{Renditions: []ninjs.Rendition{{Name: "value"}}},
			expectedErr: nil,
		},
		{
			name:        "when object has invalid renditions",
			actual:      ninjs.Item{Renditions: []ninjs.Rendition{{Name: ""}}},
			expectedErr: fmt.Errorf("error in rendition 0: field name: required"),
		},
		{
			name:        "when object has both valid and invalid renditions",
			actual:      ninjs.Item{Renditions: []ninjs.Rendition{{Name: "value"}, {Name: ""}}},
			expectedErr: fmt.Errorf("error in rendition 1: field name: required"),
		},
		{
			name:        "when object has empty renditions",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object has valid headlines",
			actual:      ninjs.Item{Headlines: []ninjs.Headline{{Value: "value"}}},
			expectedErr: nil,
		},
		{
			name:        "when object has invalid headlines",
			actual:      ninjs.Item{Headlines: []ninjs.Headline{{Value: ""}}},
			expectedErr: fmt.Errorf("error in headline 0: field value: required"),
		},
		{
			name:        "when object has both valid and invalid headlines",
			actual:      ninjs.Item{Headlines: []ninjs.Headline{{Value: "value"}, {Value: ""}}},
			expectedErr: fmt.Errorf("error in headline 1: field value: required"),
		},
		{
			name:        "when object has empty headlines",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
		{
			name:        "when object has valid descriptions",
			actual:      ninjs.Item{Descriptions: []ninjs.Description{{Value: "value"}}},
			expectedErr: nil,
		},
		{
			name:        "when object has invalid descriptions",
			actual:      ninjs.Item{Descriptions: []ninjs.Description{{Value: ""}}},
			expectedErr: fmt.Errorf("error in description 0: field value: required"),
		},
		{
			name:        "when object has both valid and invalid descriptions",
			actual:      ninjs.Item{Descriptions: []ninjs.Description{{Value: "value"}, {Value: ""}}},
			expectedErr: fmt.Errorf("error in description 1: field value: required"),
		},
		{
			name:        "when object has empty descriptions",
			actual:      ninjs.Item{},
			expectedErr: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.actual.Validate()
			DiffError(t, tc.expectedErr, err)
		})
	}
}
