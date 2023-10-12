package main

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	// import the dialect
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

func main() {
	// look up the dialect
	dialect := goqu.Dialect("postgres")

	paraName := "id"
	paraValue := "10 or 1=1"

	goqu.SetDefaultPrepared(true)

	// use dialect.From to get a dataset to build your SQL
	ds := dialect.From("test").Where(goqu.Ex{paraName: paraValue})
	sqlContent, args, err := ds.ToSQL()
	if err != nil {
		fmt.Println("An error occurred while generating the SQL", err.Error())
	} else {
		fmt.Println(sqlContent, args)
	}

	pgDb, err := sql.Open("postgres", "user=pg dbname=pg password=pg port=5433 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db := dialect.DB(pgDb)

	// "SELECT COUNT(*) FROM "user";
	if count, err := db.From("Users").Count(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User count = %d\n", count)
	}

	// INSERT ... ON CONFLICT
	dataset := db.Insert("Users").
		Rows(
			goqu.Record{"id": "f02d999b-3a65-4e51-b0c3-5e8a25eaa3a2", "name": "111 Test Addr"},
			goqu.Record{"id": "f02d999b-3a65-4e51-b0c3-5e8a25eaa3a1", "name": "112 Test Addr"},
		).
		OnConflict(goqu.DoUpdate("id", goqu.Record{"name": goqu.L("NOW()")}))
	sqlContent, args, err = dataset.ToSQL()
	fmt.Println(sqlContent, args, err)

	result, err := dataset.Executor().Exec()
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	fmt.Println(err)
}
