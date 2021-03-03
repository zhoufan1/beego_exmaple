package mysql_utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func Register(url string, maxOpen, maxIdle int, alias string) {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(fmt.Errorf("init mysql error , error: %v ", err))
	}
	err = orm.RegisterDataBase(alias, "mysql", url)
	if err != nil {
		panic(fmt.Errorf("register base dbManger error , error: %v ", err))
	}
	orm.SetMaxIdleConns(alias, maxIdle)
	orm.SetMaxOpenConns(alias, maxOpen)
	logs.Info("%v mysql init success .", alias)
}
