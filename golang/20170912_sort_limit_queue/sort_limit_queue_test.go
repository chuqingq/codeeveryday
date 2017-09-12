package utils

import (
	"testing"
)

func Test_test1(t *testing.T) {
	SortedLimitQueueInit(3, func(v1, v2 interface{}) bool {
		// t.Logf("compare1: %v, %v", v1, v2)
		i1, _ := v1.(int)
		i2, _ := v2.(int)

		// t.Logf("compare2: %v, %v", i1, i2)
		return i1 <= i2
	})

	q := NewSortedLimitQueue()

	array := []int{3, 5, 6, -1, 8, 4, 0, 2, 0, 7, 0, 9, -1}
	for i := 0; i < len(array); i++ {
		t.Logf("op: %v", array[i])
		if array[i] > 0 {
			p := q.Add(array[i])
			t.Logf("Add: %v", p)
		} else if array[i] == 0 {
			p := q.PopHead()
			t.Logf("PopHead: %v", p)
		} else { // < 0
			p := q.PopTail()
			t.Logf("PopTail: %v", p)
		}

		// check queue valid
		if q.start+q.Length() > sortedLimitQueueCapacity {
			t.Fatalf("q.Length()[%v] > capacity", q.Length())
		}
		last := 0
		for j := 0; j < q.Length(); j++ {
			v := q.PeekN(j).(int)
			if v == 0 || v < last {
				t.Fatalf("q value error: %v", q)
			}
			last = v
		}
	}
}

