/*
Library to generate ordinal number suffixes in English
*/

package ordinal

func Suffix(n int) string {
	if n >= 11 && n <= 19 {
		return "th"
	} else {
		switch n % 10 {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		default:
			return "th"
		}
	}
}
