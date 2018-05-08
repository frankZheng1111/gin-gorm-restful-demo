package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Page struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {
	db, err = gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Page{})

	// 连接池最小连接数
	//
	db.DB().SetMaxIdleConns(10)
	// 连接池最大连接数
	//
	db.DB().SetMaxOpenConns(100)

	fmt.Println("Complete DB initialize")
}

func (page *Page) Destroy() error {
	err := db.Delete(&page).Error
	return err
}

func (page *Page) Save() error {
	err := db.Save(&page).Error
	return err
}

func GetAllPages() ([]Page, error) {
	var pages []Page
	err := db.Find(&pages).Error
	return pages, err
}

func GetPageById(id string) (Page, error) {
	var page Page
	err := db.Where("id = ?", id).First(&page).Error
	return page, err
}

func CreatePage(page Page) (Page, error) {
	err := db.Create(&page).Error
	return page, err
}
