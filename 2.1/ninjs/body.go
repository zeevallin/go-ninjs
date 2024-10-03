package ninjs

import (
	"fmt"
)

// Body is the content as text or with markup.
// For a simple body use an array with one struct only containing the value property.
// Role and ContentType are then undefined and it is up to the provider.
type Body struct {
	// The total character count in this body excluding figure captions. (Added in version 1.2 according to issue #27.). nar:charcount
	CharacterCount int `json:"charcount,omitempty"`

	// The IANA (Internet Assigned Numbers Authority) media type of the content of this body.
	// Used to be called MIME type.
	ContentType string `json:"contenttype,omitempty"`

	// The role of this body
	Role string `json:"role,omitempty"`

	// The body text identified with the above role and ContentType.
	Value string `json:"value"`

	// The total number of words in this body excluding figure captions. (Added in version 1.2 according to issue #27.). nar:wordcount
	WordCount string `json:"wordcount,omitempty"`
}

// Validate checks the required fields for the Body type.
func (b Body) Validate() error {
	if b.Value == "" {
		return fmt.Errorf("field value: required")
	}
	return nil
}
