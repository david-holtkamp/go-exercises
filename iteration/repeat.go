package iteration

const repeatCount = 7

func Repeat(character string, repeatcount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
