package homework

func reverse(input []int64) (result []int64) {
	for i := int64(len(input)); i > 0; {
		result = append(result, input[i-1])
		i--
	}
	return
}
