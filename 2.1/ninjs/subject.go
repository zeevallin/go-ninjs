package ninjs

// Subject describes concepts with a relationship to the content.
type Subject struct {
	// The confidence with which the metadata has been assigned.
	Confidence int `json:"confidence,omitempty"`

	// Specifies which entity (person, organisation or system) that has created or
	// last edited the property.
	Creator string `json:"creator,omitempty"`

	// An identifier for the subject as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the subject
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the subject
	Rel string `json:"rel,omitempty"`

	// The relevance of the metadata to the news content to which it is attached.
	Relevance int `json:"relevance,omitempty"`

	// The identifier of the subject as a complete URI.
	URI string `json:"uri,omitempty"`
}
