package go_data_structure

import (
	"github.com/isdamir/assert"
	"testing"
)

func TestSet(t *testing.T){
	set := NewSet()

	set.Add(1)
	set.Add("2")
	assert.Equal(t, true, set.IsExist(1))
	assert.Equal(t, true, set.IsExist("2"))
	assert.Equal(t, 2, set.Sizeof())

	set.Remove(1)
	assert.Equal(t, true, set.IsExist("2"))
	assert.Equal(t, false, set.IsExist(1))
	assert.Equal(t, 1, set.Sizeof())
}
