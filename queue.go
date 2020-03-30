
package go_data_structure

import "sync"

/*
	使用 Slice 实现 Queue 操作
	支持：
	* 创建 Queue :
	* 删除队列中的元素:  Remove(ele interface{})
	* 获取队列第一个元素: s.GetFront() interface{}
	* 获取队列末尾元素: s.GetBack() interface{}
	* 返回队列的大小: s.Sizeof() int
	* 判断队列是否为空: s.IsEmpty() bool
	* 入队 EnQueue (在队列的尾部添加元素): s.En
	* 出队 DeQueue (在队列的首部移除元素):
	* 向队列头部添加元素: s.EnQueueFront(ele interface{})
	* 队尾元素出队: s.DeQueueBack()interface{}
*/

type SliceQueue struct {
	arr []interface{}
	sync.RWMutex
}

// 创建队列
func NewSliceQueue(cap int) *SliceQueue {
	return &SliceQueue{arr: make([]interface{}, 0, cap)}
}

// 判断队列是否为空
func (q *SliceQueue) IsEmpty() bool {
	q.RLock()
	l := len(q.arr)
	q.RUnlock()

	return l == ZERO_LEN
}

// 返回队列的长度
func (q *SliceQueue) Sizeof() int {
	q.RLock()
	size := len(q.arr)
	q.RUnlock()

	return size
}

// 尾部入队
func (q *SliceQueue) EnQueue(ele interface{}) interface{}{
	q.Lock()
	q.arr = append(q.arr, ele)
	q.Unlock()

	return ele
}

// 头部入队 EnQueueFront
func (q *SliceQueue) EnQueueFront(ele interface{}) interface{}{
	arr := make([]interface{}, 1, cap(q.arr))
	arr[0] = ele

	q.Lock()
	q.arr = append(arr, q.arr...)
	q.Unlock()

	return ele
}

// 尾部出队 DeQueueBack
func (q *SliceQueue) DeQueueBack() interface{} {
	if q.IsEmpty() {
		return nil
	}

	q.Lock()
	size := len(q.arr)
	data := q.arr[size - 1]
	q.arr = q.arr[:size - 1]
	q.Unlock()

	return data
}

// 头部出队 DeQueue
func (q *SliceQueue) DeQueue() interface{}{
	if q.IsEmpty() {
		return nil
	}

	q.Lock()
	size := len(q.arr)
	data := q.arr[0]
	q.arr = q.arr[1:size]
	q.Unlock()

	return data
}

// 获取队列头部元素 Get
func (q *SliceQueue) GetQueueFront() interface{}{
	if q.IsEmpty() {
		return nil
	}

	q.RLock()
	data := q.arr[0]
	q.RUnlock()

	return data
}

// 获取队尾元素
func (q *SliceQueue) GetQueueBack() interface{}{
	if q.IsEmpty() {
		return nil
	}

	q.RLock()
	data := q.arr[len(q.arr) - 1]
	q.RUnlock()

	return data
}



