package circularrotation

func QueryArrayCircularRotation(a []int32, k int32, queries []int32) []int32 {
	k = k % int32(len(a))

	pre := a[(int32(len(a)) - k):]
	suf := a[:(int32(len(a)) - k)]
	a = append(pre, suf...)

	result := []int32{}
	for _, v := range queries {
		result = append(result, a[v])
	}
	return result
}
