package insertion

func InsertionSort(B []int) (c []int) {
	c = make([]int, len(B))
	copy(c, B)
	// from left to right
	for i := 1; i < len(c); i++ {
		key := c[i]
		lst := i
		// from right to left
		for j := i - 1; j > -1; j-- {
			if c[j] < key {
				break
			}
			c[j+1] = c[j]
			lst = j
		}
		c[lst] = key
	}
	return
}
