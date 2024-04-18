package electronicsshop

import "sort"

func Solve(keyboards []int32, drives []int32, b int32) int32 {
	return getMoneySpent(keyboards, drives, b)
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	sort.Slice(keyboards, func(i, j int) bool {
		return keyboards[i] < keyboards[j]
	})

	sort.Slice(drives, func(i, j int) bool {
		return drives[i] < drives[j]
	})

	totalSpent := int32(-1)

	for _, kPrice := range keyboards {
		for _, dPrice := range drives {
			if kPrice+dPrice > b {
				break
			}
			if kPrice+dPrice > totalSpent {
				totalSpent = kPrice + dPrice
			}
		}
	}

	return totalSpent
}
