package ninjs

import (
	"fmt"
)

// Headline for a news object.
type Headline struct {
	// The IANA (Internet Assigned Numbers Authority) media type of the content of this headline.
	// Used to be called MIME type.
	ContentType string `json:"contenttype,omitempty"`

	// The role of this headline.
	Role string `json:"role,omitempty"`

	// The headline identified with the above Role and ContentType.
	Value string `json:"value"`
}

// Validate checks the required fields for the Headline type.
func (h Headline) Validate() error {
	if h.Value == "" {
		return fmt.Errorf("field value: required")
	}
	return nil
}
