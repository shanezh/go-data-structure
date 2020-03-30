package go_data_structure

const ZERO_LEN int =  0

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}