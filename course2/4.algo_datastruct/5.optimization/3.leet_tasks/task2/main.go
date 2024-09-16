package main

func sortTheStudents(score [][]int, k int) [][]int {
	m := len(score)

	for i := 0; i < m-1; i++ {
		studentScore := score[i][k]

		for j := i + 1; j < m; j++ {
			nextStudentScore := score[j][k]
			if nextStudentScore > studentScore {
				score[i], score[j] = score[j], score[i]
				studentScore = score[i][k]
			}
		}
	}

	return score
}
