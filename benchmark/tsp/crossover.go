package main

import "math/rand"

func OrderedCrossover(parent1, parent2 []City) ([]City, []City) {
	length := len(parent1)
	child1 := make([]City, length)
	child2 := make([]City, length)

	start := rand.Intn(length)
	end := rand.Intn(length)

	if start > end {
		start, end = end, start
	}

	copy(child1[start:end], parent1[start:end])
	copy(child2[start:end], parent2[start:end])

	fillChild := func(child, parent []City) {
		childIndex := end % length
		for _, city := range parent {
			if !contains(child[start:end], city) {
				child[childIndex] = city
				childIndex = (childIndex + 1) % length
			}
		}
	}

	fillChild(child1, parent2)
	fillChild(child2, parent1)

	return child1, child2
}

func contains(slice []City, city City) bool {
	for _, c := range slice {
		if c == city {
			return true
		}
	}
	return false
}
