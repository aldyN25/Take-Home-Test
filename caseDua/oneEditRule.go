package casedua

func OneEditRule(s1, s2 string) bool {
	if len(s1) == len(s2) {
		return Replace(s1, s2)
	} else if len(s1)+1 == len(s2) {
		return Insert(s1, s2)
	} else if len(s1)-1 == len(s2) {
		return Insert(s2, s1)
	}
	return false
}

func Replace(s1, s2 string) bool {
	foundDifference := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}

func Insert(s1, s2 string) bool {
	index1 := 0
	index2 := 0
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}
