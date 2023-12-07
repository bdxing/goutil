package goutil

// SliceInInt slice 里是否存在 element
func SliceInInt(element int, slice []int) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint slice 里是否存在 element
func SliceInUint(element uint, slice []uint) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt8 slice 里是否存在 element
func SliceInInt8(element int8, slice []int8) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint8 slice 里是否存在 element
func SliceInUint8(element uint8, slice []uint8) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt32 slice 里是否存在 element
func SliceInInt32(element int32, slice []int32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint32 slice 里是否存在 element
func SliceInUint32(element uint32, slice []uint32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInInt64 slice 里是否存在 element
func SliceInInt64(element int64, slice []int64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInUint64 slice 里是否存在 element
func SliceInUint64(element uint64, slice []uint64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInFloat32 slice 里是否存在 element
func SliceInFloat32(element float32, slice []float32) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInFloat64 slice 里是否存在 element
func SliceInFloat64(element float64, slice []float64) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}

// SliceInStr slice 里是否存在 element
func SliceInStr(element string, slice []string) bool {
	for k := range slice {
		if slice[k] == element {
			return true
		}
	}
	return false
}
