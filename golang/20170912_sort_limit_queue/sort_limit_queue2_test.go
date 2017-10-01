package utils

import (
	"testing"
)

func Test_test1(t *testing.T) {
	SortedLimitQueueInit(10)

	q := NewSortedLimitQueue()

	array := []int{3, 5, 6, -1, 8, 4, 0, 2, 0, 7, 0, 9, -1, 20, 11, 14, 13}
	// array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < len(array); i++ {
		t.Logf("op: %v", array[i])
		if array[i] > 0 {
			p := q.Add(array[i], array[i])
			t.Logf("Add: %v", p)
		} else if array[i] == 0 {
			p := q.PopHead()
			t.Logf("PopHead: %v", p)
		} else { // < 0
			p := q.PopTail()
			t.Logf("PopTail: %v", p)
		}

		t.Logf("q: %v", q)

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

	t.Fatalf("q: %v", q)
}

