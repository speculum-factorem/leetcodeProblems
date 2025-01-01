package main

func maxScore(s string) int {

	maxScore := 0
	count0 := 0
	count1 := 0

	for _, char := range s {
		if char == '1' {
			count1++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			count0++
		} else {
			count1--
		}

		if maxScore < (count0 + count1) {
			maxScore = count0 + count1
		}
	}

	return maxScore
}
