package util

func MultiInterSection(d [][]int) []int {
	switch len(d) {
	case 0:
		return nil
	case 1:
		return d[0]
	case 2:
		return TwoInterSection(d[0], d[1])
	}
	s1 := MultiInterSection(d[0 : len(d)/2])
	s2 := MultiInterSection(d[len(d)/2:])
	return TwoInterSection(s1, s2)
}

func TwoInterSection(s1, s2 []int) []int {
	if s1[0] > s2[len(s2)-1] || s2[0] > s1[len(s1)-1] {
		return nil
	}

	ret := make([]int, 0, min(len(s1), len(s2)))
	for i, j := 0, 0; i < len(s1) && j < len(s2); i++ {
		for j < len(s2) && s1[i] > s2[j] {
			j++
		}
		if j == len(s2) {
			break
		}
		if s1[i] == s2[j] {
			ret = append(ret, s2[j])
			j++
		}
		for i < len(s1)-1 && s1[i] == s1[i+1] {
			i++
		}
	}
	return ret
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
