package ninjs

import (
	"fmt"
	"time"
)

// Item is the root element of ninjs as a type to make it possible to combine schema parts.
type Item struct {
	// Alternative identifiers assigned to the content. Each alternative id can have a role and a value. nar:altId issue #3.
	AltIDs []AltID `json:"altids"`

	// An array of news objects which are associated with this news object.
	Associations []Item `json:"associations"`

	// An array of body objects with the content as text or with markup.
	// For a simple body use an array with one object only containing the value property.
	// Role and ContentType are then undefined and it is up to the provider.
	Bodies []Body `json:"bodies"`

	// A natural-language statement about the creator (author, photographer etc.) of the content.
	//
	// nar:by
	By string `json:"by"`

	// The date and time when the content of this ninjs object was originally created.
	// For example an old photo that is now handled as a ninjs object.
	//
	// nar:contentCreated
	ContentCreated time.Time `json:"contentcreated"`

	// The person or organisation claiming the intellectual property for the content.
	//
	// nar:copyrightHolder
	CopyrightHolder string `json:"copyrightholder"`

	// Any necessary copyright notice for claiming the intellectual property for the content.
	//
	// nar:copyrightNotice
	CopyrightNotice string `json:"copyrightnotice"`

	// An array of one or more descriptions of the ninjs object. See also ednote for information from provider to receiver.
	// Descriptions are seen as metadata.
	// For a simple description use an array with one object only containing the value property.
	// Role and ContentType are then undefined and it is up to the provider.
	Descriptions []Description `json:"descriptions"`

	// A note that is intended to be read by internal staff at the receiving organisation, but not intended to be published.
	//
	// (Added in version 1.2 from issue #6.).
	//
	// (Consider using this before using the descriptions array.)
	//
	// ednote: nar:edNote
	EDNote string `json:"ednote"`

	// The date and time before which all versions of the news object are embargoed.
	// If absent, this object is not embargoed. nar:embargoed
	Embargoed time.Time `json:"embargoed"`

	// An array of objects describing something which happens in a planned or unplanned manner. nar:?
	Events []Event `json:"events"`

	// The date and time after which the Item is no longer considered editorially relevant by its provider.
	//
	// (Added in 2.1)
	//
	// nar:expires
	Expires time.Time `json:"expires"`

	// Indicates when the first version of this ninjs object was created.
	//
	// (Added in version 1.2 from issue #5)
	//
	// nar:firstCreated
	FirstCreated time.Time `json:"firstcreated"`

	// A nature, intellectual or journalistic form of the content. nar:genre. (Added in version 1.3)
	Genres []Genre `json:"genres"`

	// An array of objects containing various types of headlines.
	// For a simple headline use an array with one object only containing the value property.
	// Role and ContentType are then undefined and it is up to the provider.
	Headlines []Headline `json:"headlines"`

	// An array of parties (person or organisation) which originated, modified, enhanced, distributed, aggregated or supplied the content or provided some information used to create or enhance the content.
	//
	// (Added in version 1.2 according to issue #15.).
	//
	// infosource: nar:infoSource
	InfoSources []InfoSource `json:"infosources"`

	// The human language used by the content.
	// The value should follow IETF BCP47.
	//
	// nar:language
	Language string `json:"language"`

	// The name of the location from which the content originates.
	//
	// nar:located
	Located string `json:"located"`

	// An array of objects describing something material, excluding persons.
	//
	// nar:subject
	Objects []Object `json:"objects"`

	// An array of objects describing administrative and functional structures which may, for example, act as a business, as a political party or not-for-profit party.
	//
	// nar:subject
	Organisations []Organisation `json:"organisations,omitempty" yaml:"organisations,omitempty" mapstructure:"organisations,omitempty"`

	// An array of objects describing individual human beings.
	//
	// nar:subject
	People []Person `json:"people,omitempty" yaml:"people,omitempty" mapstructure:"people,omitempty"`

	// An array of named locations.
	//
	// nar:subject
	Places []Place `json:"places,omitempty" yaml:"places,omitempty" mapstructure:"places,omitempty"`

	// An identifier for the structure of the news object.
	// This can be any string but we suggest something identifying the structure of the content such as 'text-only' or 'text-photo'.
	// Profiles are typically provider-specific.
	//
	// nar:profile
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty" mapstructure:"profile,omitempty"`

	// The publishing status of the news object, its value is *usable* by default.
	//
	// nar:pubStatus
	Pubstatus Pubstatus `json:"pubstatus,omitempty" yaml:"pubstatus,omitempty" mapstructure:"pubstatus,omitempty"`

	// An array of objects with different renditions of the news object.
	//
	// nar:remoteContent
	Renditions []Rendition `json:"renditions,omitempty" yaml:"renditions,omitempty" mapstructure:"renditions,omitempty"`

	// Indicates how complete this representation of a news item is.
	// No mapping to nar.
	// Specific for ninjs.
	RepresentationType RepresentationType `json:"representationtype,omitempty" yaml:"representationtype,omitempty" mapstructure:"representationtype,omitempty"`

	// Expression of rights to be applied to content. nar:rightsInfo
	RightsInfo RightsInfo `json:"rightsinfo,omitempty" yaml:"rightsinfo,omitempty" mapstructure:"rightsinfo,omitempty"`

	// A human-readable identifier for the item.
	//
	// (Added in version 1.2 from issue #4.)
	//
	// nar:slugline
	Slugline string `json:"slugline,omitempty" yaml:"slugline,omitempty" mapstructure:"slugline,omitempty"`

	// An object with information about standard, version and schema this instance is
	// valid against.
	//
	// nar:standard, nar:standardversion and xml:schema issue #43.
	//
	// (Added in version 1.3)
	Standard Standard `json:"standard,omitempty" yaml:"standard,omitempty" mapstructure:"standard,omitempty"`

	// An array of objects holding concepts with a relationship to the content.
	//
	// nar:subject
	Subjects []Subject `json:"subjects,omitempty" yaml:"subjects,omitempty" mapstructure:"subjects,omitempty"`

	// A short natural-language name for the item.
	// Title is metadata, use headlines for publishable headlines.
	//
	// (Added in version 1.2 according to issue #9).
	//
	// nar:itemMeta/title
	Title string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// An array of objects to allow links to documents about trust indicators.
	//
	// issue #44. (Added in version 1.3)
	TrustIndicators []TrustIndicator `json:"trustindicators,omitempty" yaml:"trustindicators,omitempty" mapstructure:"trustindicators,omitempty"`

	// The generic news type of this news object.
	//
	// (Value 'component' added in version 1.2 as issue #21.)
	// See: http://cv.iptc.org/newscodes/ninature/
	//
	// nar:itemClass
	Type ItemType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	// The editorial urgency of the content.
	// Values from 1 to 9. 1 represents the highest urgency, 9 the lowest.
	//
	// nar:urgency
	Urgency int `json:"urgency,omitempty" yaml:"urgency,omitempty" mapstructure:"urgency,omitempty"`

	// The global unique identifier for this news object. This is the only required property and should identify the ninjs object, not be used for links to external resources etc.
	//
	// nar:newsItem@guid
	URI string `json:"uri"`

	// A natural-language statement about the usage terms pertaining to the content.
	//
	// nar:usageTerms
	UsageTerms string `json:"usageterms"`

	// The version of the news object which is identified by the uri property.
	//
	// nar:newsItem@version
	Version string `json:"version"`

	// The date and time when this version of this ninjs object was created.
	//
	// nar:versionCreated
	VersionCreated time.Time `json:"versioncreated"`
}

// Validate checks if the Item and all children are valid.
func (o Item) Validate() error {
	if err := o.RepresentationType.Validate(); err != nil && o.RepresentationType != "" {
		return fmt.Errorf("error in representation type: %w", err)
	}
	if err := o.Pubstatus.Validate(); err != nil && o.Pubstatus != "" {
		return fmt.Errorf("error in pubstatus: %w", err)
	}
	if err := o.Type.Validate(); err != nil && o.Type != "" {
		return fmt.Errorf("error in type: %w", err)
	}
	for i, body := range o.Bodies {
		if err := body.Validate(); err != nil {
			return fmt.Errorf("error in body %d: %w", i, err)
		}
	}
	for i, rendition := range o.Renditions {
		if err := rendition.Validate(); err != nil {
			return fmt.Errorf("error in rendition %d: %w", i, err)
		}
	}
	for i, headline := range o.Headlines {
		if err := headline.Validate(); err != nil {
			return fmt.Errorf("error in headline %d: %w", i, err)
		}
	}
	for i, description := range o.Descriptions {
		if err := description.Validate(); err != nil {
			return fmt.Errorf("error in description %d: %w", i, err)
		}
	}
	return nil
}
