package bubble

func BubbleSort(B []int) (c []int) {
	c = make([]int, len(B))
	copy(c, B)
	for i := 0; i < len(c)-1; i++ {
		for j := 0; j < len(c)-1-i; j++ {
			if c[j] > c[j+1] {
				tmp := c[j]
				c[j] = c[j+1]
				c[j+1] = tmp
			}
		}
	}
	return
}
