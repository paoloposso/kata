package flatlandstations

import "fmt"

func main() {
	fmt.Println(flatlandSpaceStations(6, []int32{0, 1, 2, 3, 4, 5}))
}

func flatlandSpaceStations(n int32, c []int32) int32 {
	nearestMap := make(map[int32]*int32)

	for currentStation := int32(0); currentStation < n; currentStation++ {
		dist1, exists := nearestAhead(currentStation, c, n)
		if exists {
			nearestMap[currentStation] = &dist1
		}

		dist2, exists := nearestBehind(currentStation, c)

		if exists && (nearestMap[currentStation] == nil || dist2 < dist1) {
			nearestMap[currentStation] = &dist2
		}
	}

	return longestDistance(nearestMap)
}

func longestDistance(m map[int32]*int32) int32 {
	result := int32(0)

	for _, value := range m {
		if *value > int32(result) {
			result = *value
		}
	}

	return result
}

func contains(c []int32, station int32) bool {
	for i := 0; i < len(c); i++ {
		if c[i] == station {
			return true
		}
	}
	return false
}

func nearestAhead(station int32, c []int32, stationsCount int32) (int32, bool) {
	for current := station; current < int32(stationsCount); current++ {
		if contains(c, int32(current)) {
			return current - station, true
		}
	}
	return -1, false
}

func nearestBehind(station int32, c []int32) (int32, bool) {
	for current := station; current >= 0; current-- {
		if contains(c, int32(current)) {
			return station - current, true
		}
	}
	return -1, false
}
