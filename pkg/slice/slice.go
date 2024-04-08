package slice

func MapToSlice[K comparable, V any](mapa map[K]V) []V {
	slice := make([]V, 0, len(mapa))
	for _, v := range mapa {
		slice = append(slice, v)
	}
	return slice
}

// ToMap возращает map, где ключом является поле структуры, а значением сама структура
// Example:
// AccountGroupsMap := slice.ToMap(_accountGroups, func(ag model.AccountGroup) uint32 { return ag.ID })
func ToMap[K comparable, V any](slice []V, field func(V) K) map[K]V {
	mapBySlise := make(map[K]V, len(slice))
	for _, v := range slice {
		mapBySlise[field(v)] = v
	}
	return mapBySlise
}

func GetUniqueByField[K comparable, V any](slice []V, field func(V) K) []V {
	return MapToSlice(ToMap(slice, field))
}
