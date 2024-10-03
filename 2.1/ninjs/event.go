package ninjs

// Event describes something which happens in a planned or unplanned manner.
type Event struct {
	// An identifier for the event as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the event
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the event
	Rel string `json:"rel,omitempty"`

	// The identifier for the event as a complete uri
	URI string `json:"uri,omitempty"`
}
