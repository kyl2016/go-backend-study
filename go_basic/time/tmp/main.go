package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println(getEarlistBoardEndtimeDuration(time.Now().Unix(), []int64{time.Now().Unix() + 100}))
}

func elasticDuration(duration time.Duration, baseDuration time.Duration) time.Duration {
	return (baseDuration - duration) + time.Duration(rand.Int63n(int64(duration)*2))
}

func getEarlistBoardEndtimeDuration(current int64, ends []int64) time.Duration {
	ts := unixTimestamps(ends)
	sort.Sort(ts)
	for _, t := range ts {
		if t == 0 {
			continue
		}
		return time.Second * time.Duration(t-current)
	}
	return elasticDuration(elasticRangeOfExpiration, baseBoardCacheExpiration)
}

var (
	elasticRangeOfExpiration time.Duration = time.Duration(10) * time.Minute
	baseBoardCacheExpiration time.Duration = time.Duration(24) * time.Hour
)

type unixTimestamps []int64

func (ts unixTimestamps) Len() int {
	return len(ts)
}

func (ts unixTimestamps) Less(i, j int) bool {
	return ts[i] < ts[j]
}

func (ts unixTimestamps) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
