package movies

import (
	"math"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	result := 0

	for x := i; x <= j; x++ {
		reversedStr := reverseString(strconv.Itoa(int(x)))
		if revx, err := strconv.Atoi(reversedStr); err == nil {
			diff := int32(math.Abs(float64(int32(revx) - x)))

			if diff%k == 0 {
				result++
			}
		}
	}

	return int32(result)
}

func reverseString(input string) string {
	// Convert the string to a byte slice
	bytes := []byte(input)

	// Reverse the byte slice
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// Convert the reversed byte slice back to a string
	reversed := string(bytes)

	return reversed
}
