package database

import "github.com/jinzhu/gorm"
type DB struct {
	*gorm.DB
	migrations map [string]*Migration
	migrationsSeq []string
}

type testModel struct {
	gorm.Model
	Str  string         `gorm:"type:varchar(255);"`
	Subs []testSubModel `gorm:"foreignkey:ParentID;"`
}

type testSubModel struct {
	gorm.Model
	Str      string    `gorm:"type:varchar(255);"`
	Parent   testModel `gorm:"foreignkey:ParentID;"`
	ParentID uint
}

var serverDB DB

type Tx struct {
	DB
}

func init()  {
	var db *gorm.DB
	var err error
	db,err = gorm.Open("postgres", "postgres://postgres:@127.0.0.1:5432/splat?sslmode=disable")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	db.Create(&TestModel{Str:"asdasdf", Subs:[]testSubModel{{Str: "sub1"}, {Str: "sub2"}}})

}
