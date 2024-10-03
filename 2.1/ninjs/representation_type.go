package ninjs

import (
	"fmt"
	"sort"
	"strings"
)

// RepresentationType is the type of representation
type RepresentationType string

const (
	// RepresentationTypeFull represents when the representation is full.
	RepresentationTypeFull RepresentationType = "full"
	// RepresentationTypePartial represents when the representation is partial.
	RepresentationTypePartial RepresentationType = "partial"
)

// RepresentationTypeEnum is an enumeration for RepresentationType
var RepresentationTypeEnum = map[RepresentationType]struct{}{
	RepresentationTypeFull:    {},
	RepresentationTypePartial: {},
}

func representationTypeKeys() []string {
	keys := []string{}
	for key := range RepresentationTypeEnum {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)
	return keys
}

func (e RepresentationType) String() string {
	return string(e)
}

// Validate checks if the value is valid.
func (e RepresentationType) Validate() error {
	if _, ok := RepresentationTypeEnum[e]; !ok {
		return fmt.Errorf("invalid value (expected one of %s): %s", strings.Join(representationTypeKeys(), ", "), e)
	}
	return nil
}
