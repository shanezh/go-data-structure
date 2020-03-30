
package go_data_structure

import (
"github.com/isdamir/assert"
"testing"
)

func TestSliceQueue(t *testing.T){
	queue := NewSliceQueue(2)

	queue.EnQueue(1)
	queue.EnQueue(2)
	assert.Equal(t,1,queue.GetQueueFront())
	assert.Equal(t,2,queue.GetQueueBack())
	assert.Equal(t,false,queue.IsEmpty())
	assert.Equal(t,2,queue.Sizeof())

	assert.Equal(t,1,queue.DeQueue())
	assert.Equal(t,1,queue.Sizeof())
}
