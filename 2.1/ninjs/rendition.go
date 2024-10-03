package ninjs

import (
	"fmt"
)

// Rendition of the content of the news object.
// (Description changed in version 1.2 according to issue #17.)
type Rendition struct {
	// A media type which applies to this rendition. nar:remoteContent@contenttype
	ContentType string `json:"contenttype,omitempty"`

	// The total time duration of the content in seconds. (Added in version 1.2. Issue #18).
	//
	// nar:remoteContent@duration
	Duration int `json:"duration,omitempty"`

	// A refinement of a generic content type (i.e. IANA media type) by a literal
	// string value. nar:remoteContent@contenttypevariant and nar:remoteContent@format
	Format string `json:"format,omitempty"`

	// For still and moving images: the height of the display area measured in pixels.
	// nar:remoteContent@height
	Height int `json:"height,omitempty"`

	// The URL for accessing the rendition as a resource. nar:remoteContent@ref
	Href string `json:"href,omitempty"`

	// The name of this object in the array of renditions. For example 'thumbnail'
	Name string `json:"name"`

	// The size of the rendition resource in bytes
	SizeInBytes int `json:"sizeinbytes,omitempty"`

	// A title for the link to the rendition resource
	Title string `json:"title,omitempty"`

	// For still and moving images: the width of the display area measured in pixels.
	// nar:remoteContent@width
	Width int `json:"width,omitempty"`
}

// Validate checks the required fields for the Rendition type.
func (r Rendition) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("field name: required")
	}
	return nil
}
