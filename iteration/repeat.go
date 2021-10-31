package iteration

func Repeat(a string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += a
	}
	return repeated
}
