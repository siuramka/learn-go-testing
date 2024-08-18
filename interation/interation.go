package interation

func Repeat(char string, count int) string {
	if len(char) != 1 {
		panic("char is not char")
	}
	sum := ""
	for i := 0; i < count; i++ {
		sum += char
	}
	return sum
}
