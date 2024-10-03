package ninjs

// Object describes something material, excluding persons.
type Object struct {
	// An identifier for the object as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the object
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the object
	Rel string `json:"rel,omitempty"`

	// The identifier for the object as a complete URI.
	URI string `json:"uri,omitempty"`
}
