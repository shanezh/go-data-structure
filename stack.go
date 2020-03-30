package go_data_structure

import "sync"

/*
	通过 Slice 实现 Stack
	支持：
	* 创建 Stack
	* 判断是否为 Empty Stack:	s.IsEmpty
	* 返回当前栈的元素数量:	s.Sizeof()
	* 返回栈顶元素: s.GetBack()
	* 弹出栈顶元素: s.Pop()
	* 向栈顶压入元素: s.Push(ele)
*/

type SliceStack struct {
	arr []interface{}
	sync.RWMutex
}

func NewSliceStack(cap int) *SliceStack {
	return &SliceStack{arr:make([]interface{}, 0, cap)}
}

//判断是否为空
func (s *SliceStack) IsEmpty() bool {
	s.RLock()
	l := len(s.arr)
	s.RUnlock()

	return l == 0
}

//返回大小
func (s *SliceStack) Sizeof() int {
	s.RLock()
	l := len(s.arr)
	s.RUnlock()

	return l
}

//返回栈顶元素
func (s *SliceStack)GetBack()interface{}{
	if s.IsEmpty() {
		return nil
	}

	s.RLock()
	ret := s.arr[len(s.arr) - 1 ]
	s.RUnlock()

	return ret
}

//弹出栈元素
func (s *SliceStack)Pop()interface{}{
	if s.IsEmpty() {
		return nil
	}

	s.Lock()
	ret := s.arr[len(s.arr)-1]    // 返回栈顶元素
	s.arr = s.arr[:len(s.arr)-1]     // 重新切片
	s.Unlock()

	return ret
}

//Push栈元素
func (s *SliceStack) Push(ele interface{}) interface{}{
	s.Lock()
	s.arr = append(s.arr,ele)
	s.Unlock()
	return ele
}


/*
	通过 Link 实现 Stack
	支持：
	* 创建 Stack
	* 判断是否为 Empty Stack:	s.IsEmpty
	* 返回当前栈的元素数量:	s.Sizeof()
	* 返回栈顶元素: s.GetBack()
	* 弹出栈顶元素: s.Pop()
	* 向栈顶压入元素: s.Push(ele)
*/


type LinkStack struct {
	size int
	head *LNode
	sync.RWMutex
}

// 创建 LinkStack
func NewLinkStack()*LinkStack{
	return &LinkStack{size:0, head:&LNode{}}
}

func (s *LinkStack) IsEmpty() bool {
	s.RLock()
	size := s.size
	s.RUnlock()

	return size == 0
}

func (s *LinkStack) Sizeof() int {
	s.RLock()
	size := s.size
	s.RUnlock()

	return size
}

// 返回栈顶元素
func (s *LinkStack) GetBack() interface{} {
	if s.IsEmpty(){
		return nil
	}

	s.RLock()
	data := s.head.Next.Data
	s.RUnlock()

	return data
}

// 弹出栈顶元素
func (s *LinkStack) Pop() interface{}{
	if s.IsEmpty() {
		return nil
	}

	s.Lock()

	data := s.head.Next.Data
	s.head.Next = s.head.Next.Next
	s.size--

	s.Unlock()

	return data
}

// 向栈顶压入元素
func (s *LinkStack) Push(ele interface{}) interface{}{
	s.Lock()

	node := &LNode{Data:ele,Next:s.head.Next}
	s.head.Next = node
	s.size++

	s.Unlock()

	return ele
}








