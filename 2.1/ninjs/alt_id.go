package ninjs

// AltID is the alternative identifiers assigned to the content. Each alternative id can have a role and a value. nar:altId issue #3.
type AltID struct {
	// The role of the alternative id
	Role string `json:"role,omitempty"`

	// The alternative id value
	Value string `json:"value,omitempty"`
}
