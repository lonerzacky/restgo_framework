package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

func ConnectDB() *gorm.DB {
	godotenv.Load()
	db, err := gorm.Open(os.Getenv("DB_CONNECTION"), fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE")))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
