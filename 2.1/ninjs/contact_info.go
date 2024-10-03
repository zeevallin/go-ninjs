package ninjs

// ContactInfo is a contact method for a person, place or organisation.
type ContactInfo struct {
	// The address of a person, place or organisation.
	Address ContactInfoAddress `json:"address,omitempty"`

	// If this ContactInfo struct need to be qualified with what language it is in.
	// The value should follow IETF BCP47.
	Lang string `json:"lang,omitempty"`

	// Human readable name of the contact method, like name for a web page, name of persons twitter account etc.
	Name string `json:"name,omitempty"`

	// Role refers to type and could be private, office etc.
	Role string `json:"role,omitempty"`

	// Type would be method of communication like phone, mobile, address etc.
	Type string `json:"type,omitempty"`

	// Actual phone number, email address, web url etc.
	Value string `json:"value,omitempty"`
}

// ContactInfoAddress is the address of a person, place or organisation.
type ContactInfoAddress struct {
	// A subdivision of a country part of the address.
	Area string `json:"area,omitempty"`

	// A country part of the address.
	Country string `json:"country,omitempty"`

	// An array of lines to construct an address. The order is important to construct a correct address.
	Lines []string `json:"lines,omitempty"`

	// A city/town/village etc. part of the address.
	Locality string `json:"locality,omitempty"`

	// A postal code part of the address.
	PostalCode string `json:"postalcode,omitempty"`
}
