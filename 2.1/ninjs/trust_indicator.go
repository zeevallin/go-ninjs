package ninjs

// TrustIndicator links to a document about trust indicators. issue
// #44. (Added in version 1.3)
type TrustIndicator struct {
	// The URL for accessing the trust indicator resource.
	Href string `json:"href,omitempty"`

	// The role of the trust indicator as a complete uri
	Role string `json:"role,omitempty"`

	// The title of the resource being referenced.
	Title string `json:"title,omitempty"`
}
