package main

import "sort"

type N struct {
	Names   []string
	Heights []int
}

func (n N) Len() int { return len(n.Names) }
func (n N) Swap(i, j int) {
	n.Names[i], n.Names[j] = n.Names[j], n.Names[i]
	n.Heights[i], n.Heights[j] = n.Heights[j], n.Heights[i]
}
func (n N) Less(i, j int) bool {
	return n.Heights[i] > n.Heights[j]
}

func sortPeople(names []string, heights []int) []string {
	sort.Sort(&N{
		Names:   names,
		Heights: heights,
	})

	return names
}
