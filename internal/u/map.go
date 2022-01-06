package u

import (
	"constraints"
	"fmt"
)

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

func GetUnorderedKeys[K constraints.Ordered, T any](dict *map[K]T) []K {
	keys := make([]K, 0, len(*dict))
	for key := range *dict {
		keys = append(keys, key)
	}
	return keys
}

func GetValuesOrderedByKey[K constraints.Ordered, T any](dict *map[K]T) []T {
	keys := SortSlice(GetUnorderedKeys(dict))
	values := make([]T, 0, len(*dict))
	for _, key := range keys {
		values = append(values, GetOrPanic(dict, key))
	}
	return values
}
