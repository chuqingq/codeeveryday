package utils

// 新需求
// * 创建时指定大小
// * add时挤掉
// * 遍历
// * peektail

var sortedLimitQueueCapacity int

func SortedLimitQueueInit(capacity int) {
	sortedLimitQueueCapacity = capacity
}

type SortedLimitQueue struct {
	nodes  []node
	start  int
	length int
	// compare func(value1, value2 interface{}) bool // <=
}

type node struct {
	score int
	value interface{}
}

func NewSortedLimitQueue() *SortedLimitQueue {
	return &SortedLimitQueue{
		nodes: make([]node, sortedLimitQueueCapacity, sortedLimitQueueCapacity),
	}
}

func (s *SortedLimitQueue) insert(score int, value interface{}) {
	i := 0 // value应该排在第几个前面
	for ; i < s.length; i++ {
		if s.nodes[s.start+i].score > score {
			break
		}
	}

	if i > 0 && s.start > 0 {
		copy(s.nodes[0:], s.nodes[s.start:s.start+i])
	}

	if s.length > i && s.start != 1 {
		copy(s.nodes[i+1:], s.nodes[s.start+i:s.start+i+s.length-i+1])
	}

	s.nodes[i].score = score
	s.nodes[i].value = value

	s.start = 0
	s.length++
}

func (s *SortedLimitQueue) Add(score int, value interface{}) interface{} {
	if s.length < sortedLimitQueueCapacity {
		s.insert(score, value)
		return nil
	}

	// s.length == sortedLimitQueueCapacity
	if s.nodes[s.start+s.length-1].score < score {
		return value
	} else {
		tail := s.PopTail()
		s.insert(score, value)
		return tail
	}
}

func (s *SortedLimitQueue) PeekHead() interface{} {
	if s.length > 0 {
		return s.nodes[s.start].value
	}

	return nil
}

func (s *SortedLimitQueue) PopHead() interface{} {
	if s.length > 0 {
		res := s.nodes[s.start].value
		s.nodes[s.start].score = -1
		s.nodes[s.start].value = nil
		s.start++
		s.length--
		return res
	}

	return nil
}

func (s *SortedLimitQueue) PopTail() interface{} {
	if s.length > 0 {
		res := s.nodes[s.start+s.length-1].value
		s.nodes[s.start+s.length-1].score = -1
		s.nodes[s.start+s.length-1].value = nil
		s.length--
		return res
	}

	return nil
}

func (s *SortedLimitQueue) PeekTail() interface{} {
	if s.length > 0 {
		return s.nodes[s.start+s.length-1].value
	}

	return nil
}

func (s *SortedLimitQueue) PeekN(i int) interface{} {
	if s.Length() > i {
		return s.nodes[s.start+i].value
	}
	return nil
}

func (s *SortedLimitQueue) Length() int {
	return s.length
}

