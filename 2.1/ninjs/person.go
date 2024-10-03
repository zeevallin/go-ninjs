package ninjs

// Person describes an individual human beings.
type Person struct {
	// ContactInfo corresponds to the JSON schema field "contactinfo".
	ContactInfo []ContactInfo `json:"contactinfo,omitempty"`

	// An identifier for the person as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of a person
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the person
	Rel string `json:"rel,omitempty"`

	// The identifier for the person as a complete uri with the code.
	URI string `json:"uri,omitempty"`
}
