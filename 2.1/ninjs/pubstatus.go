package ninjs

import (
	"fmt"
	"sort"
	"strings"
)

// Pubstatus is the current publishing status of the news object
type Pubstatus string

const (
	// PubstatusCanceled represents when the content is canceled.
	PubstatusCanceled Pubstatus = "canceled"
	// PubstatusUsable represents when the content is usable.
	PubstatusUsable Pubstatus = "usable"
	// PubstatusWithheld represents when the content is withheld.
	PubstatusWithheld Pubstatus = "withheld"
)

// PubstatusEnum is an enumeration for Pubstatus
var PubstatusEnum = map[Pubstatus]struct{}{
	PubstatusCanceled: {},
	PubstatusUsable:   {},
	PubstatusWithheld: {},
}

func pubstatusKeys() []string {
	keys := []string{}
	for key := range PubstatusEnum {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)
	return keys
}

func (e Pubstatus) String() string {
	return string(e)
}

// Validate checks if the value is valid.
func (e Pubstatus) Validate() error {
	if _, ok := PubstatusEnum[e]; !ok {
		keys := make([]string, 0, len(PubstatusEnum))
		for k := range PubstatusEnum {
			keys = append(keys, string(k))
		}
		return fmt.Errorf("invalid value (expected one of %s): %s", strings.Join(pubstatusKeys(), ", "), e)
	}
	return nil
}
