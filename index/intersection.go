package main

import "fmt"

type MicroServiceInstances []*MicroServiceInstance

func (m *MicroServiceInstance) Less(b *MicroServiceInstance) bool { return m.InstanceID < b.InstanceID }
func (m *MicroServiceInstance) Equal(b *MicroServiceInstance) bool {
	return m.InstanceID == b.InstanceID
}

func MultiInterSection(d []MicroServiceInstances) MicroServiceInstances {
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

func TwoInterSection(s1, s2 MicroServiceInstances) MicroServiceInstances {
	if s2[len(s2)-1].Less(s1[0]) || s1[len(s1)-1].Less(s2[0]) {
		return nil
	}
	fmt.Println(s1[0], s2[0])
	ret := make(MicroServiceInstances, 0, min(len(s1), len(s2)))
	for i, j := 0, 0; i < len(s1) && j < len(s2); i++ {
		for j < len(s2) && s2[j].Less(s1[i]) {
			j++
		}
		if j == len(s2) {
			break
		}
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

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
