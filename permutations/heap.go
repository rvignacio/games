package permutations

func swap(values []interface{}, i, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}

func Generate(data ...interface{}) (results [][]interface{}) {

	results = append(results, data)
	result := make([]interface{}, len(data))
	copy(result, data)

	counter := make([]int, len(data))

	for i := 0; i < len(result); {
		if counter[i] < i {
			if i%2 == 0 {
				swap(result, 0, i)
			} else {
				swap(result, counter[i], i)
			}

			counter[i] += 1

			i = 0

			tmp := make([]interface{}, len(result))
			copy(tmp, result)
			results = append(results, result)
		} else {
			counter[i] = 0
			i++
		}
	}

	return
}
