package service

import (
	"errors"
	"fmt"
	"im/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:123456@(127.0.0.1:3306)/chat?charset=utf-8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.User))
	fmt.Println("init data base ok")
}
