package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"time"

	//"github.com/astaxie/beego/orm"
)

type DB struct {
	db *gorm.DB
}

var (
	db *gorm.DB
)

func init() {
	//logs.Info("hello init function")
	var err error
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed " + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{})
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@qq.com",
			Pwd:    "123",
			Avatar: "/static/images/info-img.png",
			Role:   0,
		})
	}

}

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
