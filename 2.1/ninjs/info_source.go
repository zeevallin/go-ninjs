package ninjs

// InfoSource is a party (person or organisation) which originated, modified, enhanced, distributed, aggregated or supplied the content or provided some information used to create or enhance the content.
// (Added in version 1.2 according to issue #15.)
type InfoSource struct {
	// ContactInfo corresponds to the JSON schema field "contactinfo".
	ContactInfo []ContactInfo `json:"contactinfo,omitempty"`

	// An identifier for the InfoSource as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the InfoSource
	Name string `json:"name,omitempty"`

	// The role the InfoSource in relationship to the content as a uri.
	Role string `json:"role,omitempty"`

	// The identifier of the InfoSource as a complete uri
	URI string `json:"uri,omitempty"`
}
