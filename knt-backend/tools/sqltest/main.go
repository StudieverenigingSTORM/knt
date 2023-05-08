package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	rows, _ := db.Query("SELECT test FROM test")

	for rows.Next() {
		var test int
		rows.Scan(&test)
		println(test)
	}
	rows.Close()
	st, err := db.Prepare("INSERT INTO test (test) VALUES (3)")
	if err != nil {
		panic(err)
	}
	_, err = st.Exec()
	if err != nil {
		panic(err)
	}

	rows, _ = db.Query("SELECT test FROM test")

	for rows.Next() {
		var test int
		rows.Scan(&test)
		println(test)
	}

}
