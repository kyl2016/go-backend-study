package bulk

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
)

func TestBulkInsertLotteryDetail(t *testing.T) {
	fmt.Println("|insert|elaspsed|")
	fmt.Println("|--|--|")
	counts := []int{1, 10, 50, 100, 200, 500, 1000, 2000, 3000, 10000}

	wg := sync.WaitGroup{}
	wg.Add(2)

	f := func(index int, db *pg.DB) {
		fmt.Println(index)
		for i := 0; i < 10; i++ {
			for _, count := range counts {
				insertBulk(db, count, "AppID222222")
			}
		}
		wg.Done()
	}

	db1 := connectDB()
	go f(1, db1)
	db2 := connectDB()
	go f(2, db2)

	wg.Wait()
}

func insertBulk(db *pg.DB, count int, appID string) {
	sqlFormat := "insert into lottery_detail (user_id, app_id, activity_name, reward_type, reward_count, draw_time) values %s"
	values := ""

	for i := 0; i < count; i++ {
		values += fmt.Sprintf("('UserTestNumber%d', '%s', 'ActivityName00001', 1, 100, '%s'),", i, appID, time.Now().Format(time.RFC3339))
	}
	values = values[:len(values)-1]

	sql := fmt.Sprintf(sqlFormat, values)
	// fmt.Println(sql)

	start := time.Now()
	res, err := db.Exec(sql)
	if err != nil {
		fmt.Println(res, err)
		panic(err)
	}
	fmt.Printf("|%d|%dms|\n", count, time.Since(start).Milliseconds())
}

func connectDB() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "bytepower_test",
		Password: "bytepower_123",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
}
