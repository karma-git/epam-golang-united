package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	// sortMapKeys
	i := 0
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
		i++
	}
	sort.Ints(keys)
	// mapValues
	for _, k := range keys {
		result = append(result, input[k])
	}
	return
}
