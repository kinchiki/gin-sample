package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kinchiki/gin-sample/api/model"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@([127.0.0.1]:3306)/gin_sample?charset=utf8mb4"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
