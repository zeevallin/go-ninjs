package ninjs

import "encoding/json"

// Place describes a location.
type Place struct {
	// ContactInfo corresponds to the JSON schema field "contactinfo".
	ContactInfo []ContactInfo `json:"contactinfo,omitempty"`

	// An identifier for the place as a free-text string.
	Literal string `json:"literal,omitempty"`

	// The name of the place
	Name string `json:"name,omitempty"`

	// The relationship of the content of the news object to the place
	Rel string `json:"rel,omitempty"`

	// The GeoJSON representation of the place as a RAW message.
	//
	// schema: https://geojson.org/schema/GeoJSON.json
	GeoJSON json.RawMessage `json:"geojson,omitempty"`

	// The identifier for the place as a complete uri
	URI string `json:"uri,omitempty"`
}
