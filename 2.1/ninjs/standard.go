package ninjs

// Standard is information about standard, version and schema this instance is valid against.
// nar:standard, nar:standardversion and xml:schema issue #43.
// (Added in version 1.3)
type Standard struct {
	// For example ninjs.
	// nar:standard
	Name string `json:"name,omitempty"`

	// The uri of the json schema to use for validation.
	Schema string `json:"schema,omitempty"`

	// For example 1.3.
	// nar:standardversion
	Version string `json:"version,omitempty"`
}
