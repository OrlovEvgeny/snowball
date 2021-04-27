package french

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 2a is the removal of Verb suffixes beginning
// with "i" in the RV region.
//
func step2a(word *snowballword.SnowballWord) bool {

	// Search for the longest among the following suffixes
	// in RV and if found, delete if preceded by a non-vowel.

	suffix, suffixRunesSize := word.FirstSuffixIn(word.RVstart, len(word.RS),
		"issantes", "issaIent", "issions", "issants", "issante",
		"iraIent", "issons", "issiez", "issent", "issant", "issait",
		"issais", "irions", "issez", "isses", "iront", "irons", "iriez",
		"irent", "irait", "irais", "îtes", "îmes", "isse", "irez",
		"iras", "irai", "ira", "ies", "ît", "it", "is", "ir", "ie", "i",
	)
	if suffix != "" {
		idx := len(word.RS) - suffixRunesSize - 1
		if idx >= 0 && word.FitsInRV(suffixRunesSize+1) && isLowerVowel(word.RS[idx]) == false {
			word.RemoveLastNRunes(suffixRunesSize)
			return true
		}
	}
	return false
}
