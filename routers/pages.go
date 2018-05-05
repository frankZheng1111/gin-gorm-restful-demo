package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "gin-gorm-restful-demo/models"
)

var err error

func initPagesRouters(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", getPages)
	routerGroup.GET("/:id", getPage)
	routerGroup.POST("/", createPage)
	// routerGroup.PUT("/:id", updatePage)
	// routerGroup.DELETE("/:id", deletePage)
}

// func updatePage(c *gin.Context) {
// 	var page Page
// 	id := c.Params.ByName("id")
// 	if err := db.Where("id = ?", id).First(&page).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	}
// 	c.BindJSON(&page)
// 	db.Save(&page)
// 	c.JSON(200, page)
// }
//
func createPage(c *gin.Context) {
	var page db.Page
  var err error
	c.BindJSON(&page)
	if page, err = db.CreatePage(page); err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, page)
	}
}

func getPage(c *gin.Context) {
	id := c.Params.ByName("id")
	if page, err := db.GetPageById(id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, page)
	}
}

func getPages(c *gin.Context) {
	if pages, err := db.GetAllPages(); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pages)
	}
}
//
// func deletePage(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var page Page
// 	d := db.Where("id = ?", id).Delete(&page)
// 	fmt.Println(d)
// 	c.JSON(200, gin.H{"id#" + id: "deleted"})
// }
