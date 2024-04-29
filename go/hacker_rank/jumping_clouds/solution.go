package jumpingclouds

func jumpingOnClouds(c []int32, k int32) int32 {
	curr := 0
	energy := 100

	for {
		curr = ((curr + int(k)) % len(c))
		energy--
		if c[curr] == 1 {
			energy -= 2
		}

		if curr == 0 {
			break
		}
	}

	return int32(energy)
}
