package snowballword

import "testing"

func Benchmark_FirstSuffixIfIn(b *testing.B) {
	var testCases = []struct {
		input    string
		startPos int
		endPos   int
		suffixes []string
		suffix   string
	}{
		{"firehose", 0, 6, []string{"x", "fi"}, ""},
		{"firehose", 0, 6, []string{"x", "eho", "fi"}, "eho"},
		{"firehose", 0, 4, []string{"re", "se"}, "re"},
		{"firehose", 0, 4, []string{"se", "xfirehose"}, ""},
		{"firehose", 0, 4, []string{"fire", "xxx"}, "fire"},
		{"firehose", 1, 5, []string{"fire", "xxx"}, ""},
		// The follwoing tests shows how FirstSuffixIfIn works. It
		// first checks for the matching suffix and only then checks
		// to see if it is starts at or before startPos.  This
		// is the behavior desired for many stemming steps but
		// is somewhat counterintuitive.
		{"firehose", 1, 5, []string{"fireh", "ireh", "h"}, ""},
		{"firehose", 1, 5, []string{"ireh", "fireh", "h"}, "ireh"},
	}

	b.ResetTimer()
	b.ReportAllocs()
	w := New("firehose")
	for n := 0; n < b.N; n++ {
		for caseN, tc := range testCases {
			suffix, _ := w.FirstSuffixIfIn(tc.startPos, tc.endPos, tc.suffixes...)
			if suffix != tc.suffix {
				b.Errorf("case #%d, Expected \"{%v}\" but got \"{%v}\"", caseN, tc.suffix, suffix)
			}
		}
	}
}



