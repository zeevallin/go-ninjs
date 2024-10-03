package ninjs

// Genre is a nature, intellectual or journalistic form of the content. nar:genre. (Added in version 1.3)
type Genre struct {
	// An identifier for the genre as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the genre
	Name string `json:"name"`

	// The identifier of the genre as a complete URI
	URI string `json:"uri"`
}
