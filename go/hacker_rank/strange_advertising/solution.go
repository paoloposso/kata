package strangeadvertising

func viralAdvertising(n int32) int32 {
	cumulative := 0
	shared := 5

	for i := 0; i < int(n); i++ {
		liked := shared / 2
		shared = liked * 3
		cumulative += liked
	}

	return int32(cumulative)
}
