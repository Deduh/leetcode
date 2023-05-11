package main

func isValid(s string) bool {
	chars := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	temp := []rune{}
	for _, char := range s {
		if openChar, ok := chars[char]; ok {
			temp = append(temp, openChar)
		} else if len(temp) == 0 || char != temp[len(temp)-1] {
			return false
		} else {
			temp = temp[:len(temp)-1]
		}
	}

	return len(temp) == 0
}
