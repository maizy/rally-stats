package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValuesOrderedByKey(t *testing.T) {
	type MyStruct struct {
		Value string
	}
	dict := map[int]MyStruct{}
	dict[2] = MyStruct{"2"}
	dict[1] = MyStruct{"1"}
	dict[99] = MyStruct{"99"}

	got := GetValuesOrderedByKey(&dict)

	assert.EqualValues(t, []MyStruct{dict[1], dict[2], dict[99]}, got)
}
