package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	reverse := []rune(input)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	output = string(reverse)
	return output
}
