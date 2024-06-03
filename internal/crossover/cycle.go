package crossover

func CycleCrossover(parent1, parent2 []int) ([]int, []int) {
	n := len(parent1)
	child1 := make([]int, n)
	child2 := make([]int, n)

	for i := 0; i < n; i++ {
		child1[i] = -1
		child2[i] = -1
	}

	findIndex := func(slice []int, value int) int {
		for i, v := range slice {
			if v == value {
				return i
			}
		}
		return -1
	}

	start := 0
	for {
		if child1[start] == -1 {
			var cycle []int
			index := start
			for {
				cycle = append(cycle, index)
				value := parent1[index]
				index = findIndex(parent2, value)
				if index == start {
					break
				}
			}

			for i, idx := range cycle {
				if i%2 == 0 {
					child1[idx] = parent1[idx]
					child2[idx] = parent2[idx]
				} else {
					child1[idx] = parent2[idx]
					child2[idx] = parent1[idx]
				}
			}
		}

		start++
		if start >= n {
			break
		}
	}

	for i := 0; i < n; i++ {
		if child1[i] == -1 {
			child1[i] = parent2[i]
		}
		if child2[i] == -1 {
			child2[i] = parent1[i]
		}
	}

	return child1, child2
}
