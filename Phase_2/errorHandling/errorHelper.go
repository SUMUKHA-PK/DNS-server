package errorHandling

func ErrorHelper(val int, err error) int {
	if err != nil {
		panic(err)
	}
	return val
}
