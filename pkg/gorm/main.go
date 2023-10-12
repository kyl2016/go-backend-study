package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	url := "host=bytepower-develop-cluster.cluster-cowgwfadfddw.rds.cn-northwest-1.amazonaws.com.cn port=5432 user=bytepower_root password=rDqe0NklUtVGBNsg dbname=bytepower_stage sslmode=disable"
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	db.SingularTable(false)
	db.DB().SetMaxIdleConns(1)
	db.DB().SetMaxOpenConns(1)

	// create table
	// db.AutoMigrate(&User{})

	// batch insert multiple data
	// users := []interface{}{
	// 	&User{1, "Kitty", time.Now().AddDate(-30, 0, 0)},
	// 	&User{2, "Bob", time.Now().AddDate(-30, 0, 0)},
	// 	&User{3, "Lili", time.Now().AddDate(-30, 0, 0)},
	// }
	// BatchCreate(db, users)

	// go showPoolStats(db.DB())

	go getVersion(db.DB())
	go getVersion(db.DB())
	go getVersion(db.DB())

	select {}
}

func getVersion(db *sql.DB) {
	for i := 0; i < 100; i++ {
		var dbVersion string
		_, err := db.Exec("SELECT version()")
		fmt.Println(dbVersion, err)
	}
}

func showPoolStats(db *sql.DB) {
	for {
		// var dbVersion string
		// _, err := db.Exec("SELECT version()")
		// fmt.Println(dbVersion, err)

		fmt.Printf("%s %+v\n", time.Now().Format("03:04:05"), db.Stats())
		time.Sleep(time.Minute)
	}
}

type User struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Birthday time.Time
}
