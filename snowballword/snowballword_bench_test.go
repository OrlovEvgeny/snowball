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

	w := New("firehose")
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, tc := range testCases {
			w.FirstSuffixIfIn(tc.startPos, tc.endPos, tc.suffixes...)
		}
	}
}

/**
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
Benchmark_FirstSuffixIfIn-8      2426738               476.0 ns/op            56 B/op          4 allocs/op

*/

/**
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
Benchmark_FirstSuffixIfIn-8      1408310               728.3 ns/op             0 B/op          0 allocs/op

*/
