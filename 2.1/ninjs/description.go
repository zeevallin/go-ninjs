package ninjs

import (
	"fmt"
)

// Description of the ninjs object. See also EDNote for information from provider to receiver.
// Descriptions are seen as metadata. For a simple description use an array with one object only containing the value property.
// Role and ContentType are then undefined and it is up to the provider.
type Description struct {
	// The IANA (Internet Assigned Numbers Authority) media type of the content of
	// this description. Used to be called MIME type.
	ContentType string `json:"contenttype,omitempty"`

	// The role of this description
	Role string `json:"role,omitempty"`

	// The descriptive text identified with the above role (and ContentType).
	Value string `json:"value"`
}

// Validate checks the required fields for the Description type.
func (d Description) Validate() error {
	if d.Value == "" {
		return fmt.Errorf("field value: required")
	}
	return nil
}
