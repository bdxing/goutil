package goutil

// SliceSplitInt 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt(slice []int, n int) (ss [][]int) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint(slice []uint, n int) (ss [][]uint) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt8 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt8(slice []int8, n int) (ss [][]int8) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint8 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint8(slice []uint8, n int) (ss [][]uint8) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt32 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt32(slice []int32, n int) (ss [][]int32) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint32 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint32(slice []uint32, n int) (ss [][]uint32) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitInt64 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitInt64(slice []int64, n int) (ss [][]int64) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitUint64 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitUint64(slice []uint64, n int) (ss [][]uint64) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}

// SliceSplitStr 把 slice 元素集合拆分成由 n 个元素为一组的分组集合
func SliceSplitStr(slice []string, n int) (ss [][]string) {
	if len(slice) == 0 {
		return
	}
	if n == 0 {
		ss = append(ss, slice)
		return
	}
	for {
		if len(slice) < n {
			if len(slice) != 0 {
				ss = append(ss, slice)
			}
			break
		}
		ss = append(ss, slice[:n])
		slice = slice[n:]
	}
	return
}
