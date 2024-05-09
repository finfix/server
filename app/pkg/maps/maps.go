package maps

func ToSlice[K comparable, V any](mapa map[K]V) []V {
	slice := make([]V, 0, len(mapa))
	for _, v := range mapa {
		slice = append(slice, v)
	}
	return slice
}

func Join[K comparable, V any](leftMap, rightMap map[K]V) map[K]V {

	joinMap := make(map[K]V, len(leftMap)+len(rightMap))

	for k, v := range leftMap {
		joinMap[k] = v
	}

	for k, v := range rightMap {
		joinMap[k] = v
	}

	return joinMap
}
