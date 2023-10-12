package main

import (
	"fmt"
	"time"
)

func main() {
	// c := "10:54"
	// resetTime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 "+c+":00")
	// fmt.Println("now in systemp location:", resetTime)

	location, err := time.LoadLocation("Asia/Shanghai")
	// now := time.Now().In(location)
	// t3 := composeDatetime(now, resetTime, location)

	// fmt.Printf("%+v\n", t3)
	// fmt.Println(t3)

	// fmt.Println(now.After(t3))

	fmt.Println(err)

	t := time.Date(
		2022,
		12,
		20,
		10,
		54,
		0,
		0,
		location)

	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", location)

	fmt.Printf("%+v\n", t.In(location))
	fmt.Printf("%+v\n", t.In(location).Hour())
	fmt.Println(t.Unix())
	fmt.Println(t.String())
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Location())
}

func composeDatetime(date time.Time, t time.Time, timeLocation *time.Location) time.Time {
	fmt.Printf("%+v\n", date)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", timeLocation.String())
	fmt.Println(t.Hour(), t.Minute())

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		timeLocation)
}
