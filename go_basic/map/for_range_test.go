package main

import (
	"testing"
	"time"
)

// delete when iterator is safe
// The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next.
// If map entries that have not yet been reached are deleted during iteration, the corresponding iteration values will not be produced.
// If map entries are inserted during iteration, the behavior is implementation-dependent, but the iteration values for each entry will be produced at most once.
// If the map is nil, the number of iterations is 0.
func TestForRange(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

	go func() {
		time.Sleep(time.Second)
		delete(m, 1)
		delete(m, 2)
		delete(m, 3)
		delete(m, 4)
	}()

	for k, v := range m {
		println(k, v)
		time.Sleep(time.Second * 2)
	}
}
