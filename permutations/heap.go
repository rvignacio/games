package permutations

func swap(values []interface{}, i, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}

func Generate(data ...interface{}) (results [][]interface{}) {

	tmp := make([]interface{}, len(data))
	copy(tmp, data)
	results = append(results, tmp)

	counter := make([]int, len(data))

	for i := 0; i < len(data); {
		if counter[i] < i {
			if i%2 == 0 {
				swap(data, 0, i)
			} else {
				swap(data, counter[i], i)
			}

			counter[i] += 1

			i = 0

			tmp = make([]interface{}, len(data))
			copy(tmp, data)
			results = append(results, tmp)
		} else {
			counter[i] = 0
			i++
		}
	}

	return
}
