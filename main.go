package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
)

var db *sql.DB
var err error
var (
	name string
)

func init() {
	db, err = sql.Open("mysql",
		"root:12341324@tcp(192.168.8.9:3306)/fuhe")
	if err != nil {
		log.Fatal("链接错误")
	}

}

func SelectRowById(id int) (string, error) {
	var name string
	err = db.QueryRow("select name from wholesaler where wholesaler_id = ?", id).Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", errors.Wrap(err, "查询有误")
		}
	}
	return name, nil
}

func main() {

	name, err := SelectRowById(2)
	if err != nil {
		fmt.Printf("%+v\n", errors.Cause(err))
	}
	fmt.Printf("name = %v", name)

}
