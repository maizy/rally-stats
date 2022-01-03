package u

import (
	"fmt"
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
