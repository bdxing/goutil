package goutil

// StrFormatHumpToUnder 驼峰转下划线
func StrFormatHumpToUnder(s string) string {
	sr := []rune(s)

	var upperIndex []int
	for i := range sr {
		if !(sr[i] >= 'A' && sr[i] <= 'Z') {
			continue
		}
		if i == 0 {
			sr[i] += 0x20
		} else {
			upperIndex = append(upperIndex, i)
		}
	}

	sr = append(sr, make([]rune, len(upperIndex))...)

	for k, v := range upperIndex {
		v += k
		sr[v] += 0x20
		copy(sr[v+1:], sr[v:])
		sr[v] = '_'
	}

	return string(sr)
}

// StrFormatUnderToHump 下划线转驼峰
func StrFormatUnderToHump(s string) string {
	sr := []rune(s)

	var humpIndex []int
	for i := range sr {
		if i == 0 &&
			(sr[i] >= 'a' && sr[i] <= 'z') {
			sr[i] -= 0x20
		}
		if sr[i] == '_' &&
			len(sr) > i+1 &&
			(sr[i+1] >= 'a' && sr[i+1] <= 'z') {
			humpIndex = append(humpIndex, i)
		}
	}

	for k, v := range humpIndex {
		v -= k
		copy(sr[v:], sr[v+1:])
		sr[v] -= 0x20
	}

	return string(sr[:len(sr)-len(humpIndex)])
}
