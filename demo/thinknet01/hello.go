package main

func sort(n ...int) []int {
	if len(n) > 1 && n[0] > n[1] {
		n[0], n[1] = n[1], n[0]
	}

	if len(n) > 2 && n[1] > n[2] {
		n[1], n[2] = n[2], n[1]
	}

	return n
}