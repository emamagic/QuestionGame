package slice

func DoesExist(list []uint, value uint) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func MapFromUint64ToUint(l []uint64) []uint {
	r := make([]uint, len(l))
	for i := range l {
		r[i] = uint(l[i])
	}

	return r
}

func MapFromUintToUint64(l []uint) []uint64 {
	r := make([]uint64, len(l))
	for i := range l {
		r[i] = uint64(l[i])
	}

	return r
}