package u

import "fmt"

func GetOrElse[T any, K comparable](dict *map[K]T, key K, orElse T) T {
	if dict == nil {
		return orElse
	}
	var dictV = *dict
	if v, ok := dictV[key]; ok {
		return v
	}
	return orElse
}

func GetOrPanic[T any, K comparable](dict *map[K]T, key K) T {
	if dict == nil {
		panic("non initialized map")
	}
	var dictV = *dict
	v, ok := dictV[key]
	if !ok {
		panic(fmt.Sprintf("%v must be in map", key))
	}
	return v
}
