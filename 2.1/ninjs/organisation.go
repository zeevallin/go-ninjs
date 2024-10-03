package ninjs

// Organisation describes an administrative and functional structures which may, for example, act as a business, as a political party or not-for-profit party.
type Organisation struct {
	// ContactInfo corresponds to the JSON schema field "contactinfo".
	ContactInfo []ContactInfo `json:"contactinfo,omitempty"`

	// An identifier for the organisation as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the organisation.
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the organisation.
	Rel string `json:"rel,omitempty"`

	// Symbols used for a financial instrument linked to the organisation at a specific marketplace.
	Symbols []OrganisationSymbol `json:"symbols,omitempty"`

	// The identifier of the organisation as a complete URI.
	URI string `json:"uri,omitempty"`
}

// OrganisationSymbol is a symbol used for a financial instrument linked to the organisation at a specific marketplace.
type OrganisationSymbol struct {
	// Identifier for the marketplace which uses the ticker symbols of the ticker property
	Exchange string `json:"exchange,omitempty"`

	// Compare with hasInstrument in NewsML-G2. Same as symbol in G2.
	Symbol string `json:"symbol,omitempty"`

	// https://cv.iptc.org/newscodes/financialinstrumentsymboltype.
	// Same as type in G2.
	SymbolType string `json:"symboltype,omitempty"`

	// Ticker symbol used for the financial instrument
	Ticker string `json:"ticker,omitempty"`
}
