package pkg

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var MysqlEngine *xorm.Engine

func InitMysql() {
	// dbs := "root:root@(127.0.0.1:3306)/go0615?charset=utf8mb4"
	dbs := "jobs2333:fvnPTrW3BISf1E7a@(mysql5.sqlpub.com:3310)/go0615"
	var err error
	MysqlEngine, err = xorm.NewEngine("mysql", dbs)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
	MysqlEngine.ShowSQL(true)

	fmt.Println("Successfully Init Mysql!")
}
