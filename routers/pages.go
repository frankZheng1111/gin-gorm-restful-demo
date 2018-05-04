package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func initPagesRouters(routerGroup *gin.RouterGroup) {
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

	routerGroup.GET("/", getPages)
	routerGroup.GET("/:id", getPage)
	routerGroup.POST("/", createPage)
	routerGroup.PUT("/:id", updatePage)
	routerGroup.DELETE("/:id", deletePage)
}

func updatePage(c *gin.Context) {
	var page Page
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&page).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&page)
	db.Save(&page)
	c.JSON(200, page)
}

func createPage(c *gin.Context) {
	var page Page
	c.BindJSON(&page)
	db.Create(&page)
	c.JSON(200, page)
}

func getPage(c *gin.Context) {
	id := c.Params.ByName("id")
	var page Page
	if err := db.Where("id = ?", id).First(&page).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, page)
	}
}

func getPages(c *gin.Context) {
	var page []Page
	if err := db.Find(&page).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, page)
	}
}

func deletePage(c *gin.Context) {
	id := c.Params.ByName("id")
	var page Page
	d := db.Where("id = ?", id).Delete(&page)
	fmt.Println(d)
	c.JSON(200, gin.H{"id#" + id: "deleted"})
}
