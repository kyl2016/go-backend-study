package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/kyl2016/Play-With-Golang/utility"
)

func main() {
	// debug
	dbURL := "postgres://bytepower_root:rDqe0NklUtVGBNsg@bytepower-develop-cluster.cluster-cowgwfadfddw.rds.cn-northwest-1.amazonaws.com.cn:5432/bytepower_stage?sslmode=disable"

	cfg := DatabaseConfig{
		URL:          dbURL,
		MinIdleConns: 1,
		MaxOpenConns: 1,
	}
	utility.PanicIfNotNil(checkDBConfig(&cfg))
	service, err := InitDBService(&cfg)
	utility.PanicIfNotNil(err)

	go showPoolStats(service.DB, time.Second)

	// getVersion(service.DB)

	go getVersion(service.DB)
	go getVersion(service.DB)
	go getVersion(service.DB)

	select {}
}

func getVersion(db *pg.DB) {
	for i := 0; i < 100; i++ {
		var dbVersion string
		_, err := db.QueryOne(pg.Scan(&dbVersion), "SELECT version()")
		fmt.Println(dbVersion, err)
	}
}

func showPoolStats(db *pg.DB, interval time.Duration) {
	for {
		// var dbVersion string
		// _, err := db.QueryOne(pg.Scan(&dbVersion), "SELECT version()")
		// fmt.Println(dbVersion, err)

		fmt.Printf("%s %+v\n", time.Now().Format("03:04:05"), db.PoolStats())
		time.Sleep(interval)
	}
}
