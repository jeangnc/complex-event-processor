package util

func SliceAll(s []bool) bool {
	r := true

	for _, b := range s {
		r = r && b
	}

	return r
}

func SliceAny(s []bool) bool {
	for _, b := range s {
		if b {
			return true
		}
	}

	return false
}
