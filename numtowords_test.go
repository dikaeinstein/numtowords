package numtowords_test

import (
	"testing"

	"github.com/dikaeinstein/numtowords"
)

func TestNumToWords(t *testing.T) {
	testCases := []struct {
		input  int
		result string
	}{
		{-1, "invalid input"},
		{99999, "invalid input"},
		{0, "invalid input"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twenty one"},
		{30, "thirty"},
		{31, "thirty one"},
		{56, "fifty six"},
		{90, "ninety"},
		{91, "ninety one"},
		{99, "ninety nine"},
		{568, "five hundred and sixty eight"},
		{1000, "one thousand"},
		{2000, "two thousand"},
		{1001, "one thousand and one"},
		{1012, "one thousand and twelve"},
		{1021, "one thousand and twenty one"},
		{1121, "one thousand one hundred and twenty one"},
		{9999, "nine thousand nine hundred and ninety nine"},
	}

	for _, tC := range testCases {
		words := numtowords.Convert(tC.input)
		if words != tC.result {
			t.Errorf("want %s, got %s", tC.result, words)
		}
	}
}
