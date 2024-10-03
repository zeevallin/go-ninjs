package ninjs

// RightsInfo is an expression of rights to be applied to content.
// nar:rightsInfo
type RightsInfo struct {
	// Contains a rights expression as defined by a Rights Expression Language.
	// nar:rightsExpressionXML or nar:rightsExpressionData
	EncodedRights string `json:"encodedrights,omitempty"`

	// Identifier for the Rights Expression language used.
	// nar:@langid
	LangID string `json:"langid,omitempty"`

	// A link from the current Item to Web resource with rights related information.
	// nar:link
	LinkedRights string `json:"linkedrights,omitempty"`
}
