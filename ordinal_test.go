package ordinal

import "testing"

var suffixTests = []struct {
	given    int
	expected string
}{
	{0, "th"},
	{1, "st"},
	{2, "nd"},
	{3, "rd"},
	{4, "th"},
	{5, "th"},
	// ...
	{9, "th"},
	{10, "th"},
	{11, "th"},
	// ...
	{19, "th"},
	{20, "th"},
	{21, "st"},
	{22, "nd"},
	{23, "rd"},
	{24, "th"},
	// ...
	{100, "th"},
	{101, "st"},
	{102, "nd"},
	{103, "rd"},
	{104, "th"},
	// ...
	{150, "th"},
	{151, "st"},
	{152, "nd"},
	{153, "rd"},
	{154, "th"},
}

func TestOrdinal(t *testing.T) {
	for _, test := range suffixTests {
		actual := Suffix(test.given)
		if actual != test.expected {
			t.Errorf("ordinal.Suffix(%q) => %q, want %q", test.given, actual, test.expected)
		}
	}
}
