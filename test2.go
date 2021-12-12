package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/golang")

	err :=db.Ping()
	if err != nil{
		fmt.Println("数据库链接失败")
	}
	defer db.Close()

	var name string
	row := db.QueryRow("SELECT name FROM users WHERE id = 2;")
	err1 := row.Scan(&name)
	if err1 != nil && err != sql.ErrNoRows {
		// log the error
	}
}