package dao

import (
	"database/sql"
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
		"root:root@tcp(127.0.0.1:3306)/dinghuo3")
	if err != nil {
		log.Fatal("链接错误")
	}

}

func GetNameById(id int) (string, error) {
	var name string
	err = db.QueryRow("select name from wholesaler where wholesaler_id = ?", id).Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrap(err, "数据不存在")
		} else {
			return "", errors.Wrap(err, "查询有误")
		}
	}
	return name, nil
}
