package spanish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 3 is the removal of residual suffixes.
//
func step3(word *snowballword.SnowballWord) bool {
	suffix, suffixRunesSize := word.FirstSuffixIfIn(word.RVstart, len(word.RS),
		"os", "a", "o", "á", "í", "ó", "e", "é",
	)

	// No suffix found, nothing to do.
	//
	if suffix == "" {
		return false
	}

	// Remove all these suffixes
	word.RemoveLastNRunes(suffixRunesSize)

	if suffix == "e" || suffix == "é" {

		// If preceded by gu with the u in RV delete the u
		//
		guSuffix, _ := word.FirstSuffix("gu")
		if guSuffix != "" {
			word.RemoveLastNRunes(1)
		}
	}
	return true
}
