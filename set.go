package go_data_structure

import "sync"

/*
	使用 map 实现 Set 结构
	支持：
	* 添加 key : s.Add(ele interface{})
	* 返回 Set 的大小: s.Sizeof() int {}
	* 判断 Key 是否存在： s.Exist(ele interface{}) bool {}
	* 删除 Key : s.Remove(ele interface{}){}
	* 清空 Set : s.Clear()
	* 判断 Set 集合是否为空：s.IsEmpty()
	*
*/

type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

// 创建 Set集合
func NewSet() *Set {
	return &Set{m:map[interface{}]struct{}{}}
}

// 判断集合是否为空
func (s *Set) IsEmpty() bool {
	s.RLock()
	l := len(s.m)
	s.RUnlock()

	return l == ZERO_LEN
}

// 返回集合的大小
func (s *Set) Sizeof() int{
	s.RLock()
	size := len(s.m)
	s.RUnlock()

	return size
}

// 向集合添加元素
func (s *Set) Add(k interface{}) interface{}{
	s.Lock()
	s.m[k] = struct{}{}
	s.Unlock()

	return k
}

// 删除集合中的元素
func (s *Set) Remove(k interface{}) interface{}{
	if s.IsEmpty() {
		return nil
	}

	s.Lock()
	delete(s.m, k)
	s.Unlock()

	return k
}

// 清空集合
func (s *Set) Clear(){
	s.Lock()
	s.m = map[interface{}]struct{}{}
	s.Unlock()
}

// 判断 key 是否存在
func (s *Set) IsExist(k interface{}) bool{
	if s.IsEmpty() {
		return false
	}
	s.RLock()
	_,ok := s.m[k]
	s.RUnlock()

	return ok
}