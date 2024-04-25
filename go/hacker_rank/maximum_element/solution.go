package maximumelement

import (
	"strconv"
	"strings"
)

func getMax(operations []string) []int32 {
	q := []int32{}
	result := []int32{}

	for _, v := range operations {
		operationSet := strings.Split(v, " ")

		operation := operationSet[0]

		if operation == "2" {
			q = q[:(len(q) - 1)]
		} else if operation == "1" {
			val, _ := strconv.ParseInt(operationSet[1], 10, 32)
			q = append(q, int32(val))
		} else {
			max := int32(0)
			for _, v := range q {
				if v > max {
					max = v
				}
			}
			result = append(result, max)
		}
	}
	return result
}
