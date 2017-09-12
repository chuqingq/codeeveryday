package utils

// 新需求
// * 创建时指定大小
// * add时挤掉
// * 遍历
// * peektail

const capacity = 32

type SortedLimitQueue struct {
	nodes   []interface{}
	start   int
	length  int
	compare func(value1, value2 interface{}) bool // <=
}

func NewSortedLimitQueue(compare func(value1, value2 interface{}) bool) *SortedLimitQueue {
	return &SortedLimitQueue{
		nodes:   make([]interface{}, capacity, capacity),
		compare: compare,
	}
}

func (s *SortedLimitQueue) insert(value interface{}) {
	i := 0 // value应该排在第几个前面
	for ; i < s.length; i++ {
		if !s.compare(s.nodes[s.start+i], value) {
			break
		}
	}

	if i > 0 && s.start > 0 {
		copy(s.nodes[0:], s.nodes[s.start:s.start+i])
	}

	if s.length > i {
		copy(s.nodes[i+1:], s.nodes[s.start+i:s.start+i+s.length-i+1])
	}

	s.nodes[i] = value

	s.start = 0
	s.length++
}

func (s *SortedLimitQueue) Add(value interface{}) interface{} {
	if s.length < capacity {
		s.insert(value)
		return nil
	}

	// s.length == capacity
	if s.compare(s.PeekTail(), value) {
		return value
	} else {
		tail := s.PopTail()
		s.insert(value)
		return tail
	}
}

func (s *SortedLimitQueue) PeekHead() interface{} {
	if s.length > 0 {
		return s.nodes[s.start]
	}

	return nil
}

func (s *SortedLimitQueue) PopHead() interface{} {
	if s.length > 0 {
		res := s.nodes[s.start]
		s.nodes[s.start] = nil
		s.start++
		s.length--
		return res
	}

	return nil
}

func (s *SortedLimitQueue) PopTail() interface{} {
	if s.length > 0 {
		res := s.nodes[s.start+s.length-1]
		s.nodes[s.start+s.length-1] = nil
		s.length--
		return res
	}

	return nil
}

func (s *SortedLimitQueue) PeekTail() interface{} {
	if s.length > 0 {
		return s.nodes[s.start+s.length-1]
	}

	return nil
}

func (s *SortedLimitQueue) PeekN(i int) interface{} {
	if s.Length() > i {
		return s.nodes[s.start+i]
	}
	return nil
}

func (s *SortedLimitQueue) Length() int {
	return s.length
}

