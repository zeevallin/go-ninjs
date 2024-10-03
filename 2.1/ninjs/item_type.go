package ninjs

import (
	"fmt"
	"sort"
	"strings"
)

// ItemType of the content
type ItemType string

const (
	// ItemTypeAudio represents when the content is audio.
	ItemTypeAudio ItemType = "audio"
	// ItemTypeComponent represents when the content is a component.
	ItemTypeComponent ItemType = "component"
	// ItemTypeComposite represents when the content is a composite.
	ItemTypeComposite ItemType = "composite"
	// ItemTypeGraphic represents when the content is some kind of graphic.
	ItemTypeGraphic ItemType = "graphic"
	// ItemTypePicture represents when the content is a picture.
	ItemTypePicture ItemType = "picture"
	// ItemTypeText represents when the content is text.
	ItemTypeText ItemType = "text"
	// ItemTypeVideo represents when the content is video.
	ItemTypeVideo ItemType = "video"
)

// ItemTypeEnum is an enumeration for Type
var ItemTypeEnum = map[ItemType]struct{}{
	ItemTypeAudio:     {},
	ItemTypeComponent: {},
	ItemTypeComposite: {},
	ItemTypeGraphic:   {},
	ItemTypePicture:   {},
	ItemTypeText:      {},
	ItemTypeVideo:     {},
}

func itemTypeKeys() []string {
	keys := []string{}
	for key := range ItemTypeEnum {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)
	return keys
}

func (e ItemType) String() string {
	return string(e)
}

// Validate checks if the value is valid.
func (e ItemType) Validate() error {
	if _, ok := ItemTypeEnum[e]; !ok {
		keys := make([]string, 0, len(ItemTypeEnum))
		for k := range ItemTypeEnum {
			keys = append(keys, string(k))
		}
		return fmt.Errorf("invalid value (expected one of %s): %s", strings.Join(itemTypeKeys(), ", "), e)
	}
	return nil
}
