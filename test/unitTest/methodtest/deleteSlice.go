package methodtest

func delOneFromArray(slice []int, n int) []int {
	length := len(slice)
	result := []int{}
	for i := 0; i < length; i++ {
		if n != slice[i] {
			result = append(result, slice[i])
		}
	}
	return result
}
