func isPalindrome(s string) bool {
	var r []rune
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			r = append(r, unicode.ToLower(c))
		}
	}
	for i := range len(r) / 2 {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}