package goutil

// SliceUniqueInt 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt(slice *[]int) {
	var (
		mg = map[int]struct{}{}
		i  int
		k  int
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint(slice *[]uint) {
	var (
		mg = map[uint]struct{}{}
		i  int
		k  uint
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt8 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt8(slice *[]int8) {
	var (
		mg = map[int8]struct{}{}
		i  int
		k  int8
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint8 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint8(slice *[]uint8) {
	var (
		mg = map[uint8]struct{}{}
		i  int
		k  uint8
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt32 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt32(slice *[]int32) {
	var (
		mg = map[int32]struct{}{}
		i  int
		k  int32
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint32 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint32(slice *[]uint32) {
	var (
		mg = map[uint32]struct{}{}
		i  int
		k  uint32
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueInt64 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueInt64(slice *[]int64) {
	var (
		mg = map[int64]struct{}{}
		i  int
		k  int64
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueUint64 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueUint64(slice *[]uint64) {
	var (
		mg = map[uint64]struct{}{}
		i  int
		k  uint64
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}

// SliceUniqueStr 去除 slice 里的重复元素。
// 注意：由于使用了 map 结构，所以会改变结果的顺序
func SliceUniqueStr(slice *[]string) {
	var (
		mg = map[string]struct{}{}
		i  int
		k  string
	)
	for i = range *slice {
		mg[(*slice)[i]] = struct{}{}
	}
	i = 0
	for k = range mg {
		(*slice)[i] = k
		i++
	}
	*slice = (*slice)[:i]
}
