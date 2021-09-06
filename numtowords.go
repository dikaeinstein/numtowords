package numtowords

import (
	"strconv"
	"strings"
)

// Convert takes an integer number in the range 1 - 9999
// and convert that number to words
func Convert(n int) string {
	if n <= 0 || n > 9999 {
		return "invalid input"
	}

	ones := []string{"", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine"}
	tens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tenMultiples := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty",
		"seventy", "eighty", "ninety"}
	powOfTens := []string{"hundred", "thousand"}

	nStr := strconv.Itoa(n)
	l := len(nStr)

	if l == 1 {
		return ones[n]
	}

	result := ""
	for i := 0; i < len(nStr); i++ {
		// handle the first 2 digits
		if l >= 3 {
			if nStr[i]-'0' != 0 {
				const and = "and "
				result += ones[nStr[i]-'0'] + " " + powOfTens[l-3] + " "
				if l-3 == 0 {
					result += and
				}
				if l-3 == 1 && n%1000 != 0 && nStr[i+1]-'0' == 0 {
					result += and
				}
			}
			l--
		} else {
			// handle 10 - 19
			if nStr[i] == '1' {
				sum := nStr[i] - '0' + nStr[i+1] - '0'
				result += tens[sum-1]
			} else if nStr[i] == '2' && nStr[i+1]-'0' == 0 { // handle 20
				result += tenMultiples[2]
			} else { // handle 21 - 99
				idx := nStr[i] - '0'
				if idx > 0 {
					result += tenMultiples[idx] + " "
				}
				if nStr[i+1]-'0' != 0 {
					result += ones[nStr[i+1]-'0']
				}
			}
			i++
		}
	}

	return strings.Trim(result, " ")
}
