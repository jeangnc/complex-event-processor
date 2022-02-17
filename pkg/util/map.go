package util

func MergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	newMap := CopyMap(m1)
	for k, v := range m2 {
		newMap[k] = v
	}

	return newMap
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	newMap := map[K]V{}

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}
