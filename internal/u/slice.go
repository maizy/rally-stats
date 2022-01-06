package u

import (
	"constraints"
	"fmt"
	"sort"
	"strings"
)

func JoinAsStrings[T fmt.Stringer](vals []T, div string) string {
	var result strings.Builder
	first := true
	for _, el := range vals {
		if !first {
			result.WriteString(div)
		}
		result.WriteString(el.String())
		if first {
			first = false
		}
	}
	return result.String()
}

func CopySlice[T any](vals []T) []T {
	result := make([]T, len(vals))
	copy(result, vals)
	return result
}

func SortSlice[T constraints.Ordered](vals []T) []T {
	if vals == nil {
		return nil
	}
	result := CopySlice(vals)
	sort.Slice(result, func(i, j int) bool {
		return result[i] <= result[j]
	})
	return result
}
