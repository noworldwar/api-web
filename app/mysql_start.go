package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/noworldwar/api-web/model"
	"xorm.io/core"
)

func InitMySQL() {
	model.MyDB, _ = xorm.NewEngine("mysql", "root:password@tcp(127.0.0.1:3306)/wgdb?charset=utf8mb4")
	model.MyDB.SetMapper(core.SameMapper{})

	err := model.MyDB.Ping()
	if err != nil {
		log.Fatalln("Init MySQL Error:", err)
	}

	AutoMigrate()
}

func AutoMigrate() {
	err := model.MyDB.Sync2(new(model.PlayerWallet))
	if err != nil {
		log.Fatalln("PlayerWallet Sync Error:", err)
	}
}
