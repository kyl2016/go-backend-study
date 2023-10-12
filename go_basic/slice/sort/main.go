package main

import (
	"fmt"
	"sort"
	"time"
)

type DateTimes []time.Time

func (dt DateTimes) Len() int {
	return len(dt)
}

func (dt DateTimes) Swap(i, j int) {
	dt[i], dt[j] = dt[j], dt[i]
}

func (dt DateTimes) Less(i, j int) bool {
	return dt[i].Before(dt[j])
}

func main() {
	datetimes := []time.Time{
		time.Now(),
		time.Now().Add(-time.Hour * 24),
		time.Now().Add(-time.Hour * 48),
		time.Now().Add(time.Hour),
	}
	fmt.Println(datetimes)
	sort.Slice(datetimes, func(i, j int) bool {
		return datetimes[i].Before(datetimes[j])
	})
	fmt.Println(datetimes)

	sort.Sort(DateTimes(datetimes))
	fmt.Println(datetimes)

}
