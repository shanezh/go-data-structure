package go_data_structure

import (
	"fmt"
	"github.com/isdamir/assert"
	"testing"
)

func TestSliceStack(t *testing.T){
	defer func(){
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()
	stk := NewSliceStack(2)
	stk.Push(1)
	assert.Equal(t,1,stk.GetBack().(int))
	assert.Equal(t,1,stk.Sizeof())
	assert.NotNil(t,stk.Pop())
	assert.Nil(t,stk.Pop())
	assert.Equal(t,0, stk.Sizeof())
}


func TestLinkStack(t *testing.T){
	st := NewLinkStack()

	st.Push(1)
	assert.Equal(t,1,st.GetBack().(int))
	assert.Equal(t,1,st.Sizeof())
	assert.NotNil(t,st.Pop())
	assert.Nil(t,st.Pop())
	assert.Equal(t,0,st.Sizeof())
}
