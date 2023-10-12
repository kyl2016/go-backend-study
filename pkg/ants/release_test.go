package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	pool, _ := ants.NewPool(1)

	pool.Submit(func() {
		fmt.Println("start")
		time.Sleep(time.Second * 1)
		fmt.Println("end")
	})
	pool.Release()

	fmt.Println(pool.Cap(), pool.Free())

	time.Sleep(time.Second * 2)
	fmt.Println(pool.Cap(), pool.Free())
	err := pool.Submit(func() {
		fmt.Println("new")
	})
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second)
}
