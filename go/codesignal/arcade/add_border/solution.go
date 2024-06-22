package addborder

import "fmt"

func solution(picture []string) []string {
	for i := range picture {
		picture[i] = fmt.Sprintf("*%v*", picture[i])
	}

	r := ""

	for range picture[0] {
		r += "*"
	}

	picture = append(picture, r)
	picture = append([]string{r}, picture...)

	return picture
}
