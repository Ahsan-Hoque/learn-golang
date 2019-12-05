package iteration

// Repeat single char
func Repeat(char string) (repeated string) {
	for i := 0; i < 6; i++ {
		repeated += char
	}
	return repeated
}
