package main

import (
	"context"
	"fmt"
	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"strings"
	"time"
)

// Event validate result
type Event struct {
	//tableName struct{} `pg:"event,alias:t,,discard_unknown_columns"`
	tableName struct{} `pg:"event,alias:t,discard_unknown_columns"`

	AppID      string    `pg:"app_id" json:"-"`
	Name       string    `pg:"name" json:"name"`
	State      string    `pg:"state" json:"state"`
	ErrorInfo  string    `pg:"error_info" json:"error_info,omitempty"`
	CreateTime time.Time `pg:"sys_create_time" json:"-"`
	UpdateTime time.Time `pg:"sys_update_time" json:"update_time"`
}

func main() {
	pg.SetLogger(&myLog{})
	db := getDB()
	_, err := db.Exec("select 1")
	//r, err := db.Exec("insert into event values (?, ?, ?, ?, ?, ?)", "app1", "e100", "successful", "", time.Now(), time.Now())
	//fmt.Println(r, err)

	in := "'" + strings.Join([]string{"e1", "e2"}, "','") + "'"
	fmt.Println(in)

	var events []*Event
	err = getDB().Model(&events).Where(fmt.Sprintf("app_id='%s' and name in (%s)", "111", in)).Select()
	fmt.Println(err)
	for _, e := range events {
		fmt.Println(*e)
	}
}

type HookTest struct {
	Id    int
	Value string

	beforeScan int
	afterScan  int

	afterSelect int

	beforeInsert int
	afterInsert  int

	beforeUpdate int
	afterUpdate  int

	beforeDelete int
	afterDelete  int
}

var _ pg.BeforeScanHook = (*HookTest)(nil)

func (t *HookTest) BeforeScan(c context.Context) error {
	t.beforeScan++
	return nil
}

var _ pg.AfterScanHook = (*HookTest)(nil)

func (t *HookTest) AfterScan(c context.Context) error {
	t.afterScan++
	return nil
}

var _ pg.AfterSelectHook = (*HookTest)(nil)

func (t *HookTest) AfterSelect(c context.Context) error {
	t.afterSelect++
	return nil
}

func (t *HookTest) BeforeInsert(c context.Context) (context.Context, error) {
	t.beforeInsert++
	return c, nil
}

type myLog struct {
}

func (l *myLog) Printf(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var pgConfig = map[string]string{
	"user":     "bytepower_test",
	"password": "bytepower_123",
	"database": "testdb",
	"host":     "localhost",
	"port":     "5432",
}

func getDB() orm.DB {
	pgDB := pg.Connect(&pg.Options{
		User:     pgConfig["user"],
		Password: pgConfig["password"],
		Database: pgConfig["database"],
		Addr:     pgConfig["host"] + ":" + pgConfig["port"],
	})
	if _, err := pgDB.Exec("select 1"); err != nil {
		panic(err)
	}

	return pgDB
}
