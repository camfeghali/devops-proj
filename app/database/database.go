package database

import (
	// "database/sql"
	"fmt"
	// "log"
	"hackernews-clone/graph/model"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "reflect"
	// "encoding/json"
)

var Db *gorm.DB

func InitDB() {

	// dsn := "root:supersecret@tcp(172.17.0.2:3306)/hackernoon_clone?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:supersecret@tcp(mysql-server:3306)/hackernews_clone?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	Db = db

	InitialMigration()

}

func InitialMigration() {

	Db.AutoMigrate(&model.User{}, &model.Link{})
}

