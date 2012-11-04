// Examples to use github.com/bmizerany/pq

package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/bmizerany/pq"
)

func OpenDB(conn string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Connection error: %s", err.Error())
	}

	return
}

func main() {
	args := os.Args
	if len(args) != 2 {
		conn_example := "postgres://user:pass@host:port/db"
		fmt.Println("Please provide a connection string in this format",
			conn_example)
		return
	}

	conn, err := pq.ParseURL(args[1])
	if err != nil {
		fmt.Println("Wrong format of connect string.")
		return
	}

	db, openerr := OpenDB(conn)
	if openerr != nil {
		fmt.Println("DB open error", openerr)
		return
	}

	defer db.Close()

	var i int

	row := db.QueryRow("SELECT 1")
	err_ := row.Scan(&i)

	if err_ != nil {
		fmt.Println(err_)
		return
	}

	if i != 1 {
		fmt.Println("'SELECT 1' expects 1, but we got", i)
		return
	}

	fmt.Println("Great, 'SELECT 1' returns", i)
	return
}
