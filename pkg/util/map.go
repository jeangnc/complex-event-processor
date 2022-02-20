package util

func MergeMaps[K comparable, V comparable](m1 map[K]V, m2 map[K]V) (map[K]V, []K) {
	newMap := CopyMap(m1)
	changes := make([]K, 0, 0)

	for k, v := range m2 {
		newMap[k] = v

		if m1[k] != newMap[k] {
			changes = append(changes, k)
		}
	}

	return newMap, changes
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	newMap := map[K]V{}

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}
