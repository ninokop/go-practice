package intersection

func TwoInterSection(s1, s2 MicroServiceInstances) MicroServiceInstances {
	// if s2[len(s2)-1].Less(s1[0]) || s1[len(s1)-1].Less(s2[0]) {
	// 	return nil
	// }
	fmt.Println(s1[0], s2[0])
	ret := make(MicroServiceInstances, 0, min(len(s1), len(s2)))
	for i, j := 0, 0; i < len(s1) && j < len(s2); i++ {
		for j < len(s2) && s2[j].Less(s1[j]) {
			fmt.Println(s2[j].Less(s1[j]), s2[j], s1[j])
			j++
		}
		fmt.Println(j, len(s2))
		if j == len(s2) {
			break
		}
		fmt.Println("a && b", s1[i], s2[j])
		if s1[i].Equal(s2[j]) {
			ret = append(ret, s2[j])
			j++
		}
		for i < len(s1)-1 && s1[i].Equal(s1[i+1]) {
			i++
		}
	}
	return ret
}
