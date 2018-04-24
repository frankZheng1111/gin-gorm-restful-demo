package routers

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct {
  ID uint `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
}

func InitRouters(routerVersionGroup *gin.RouterGroup) {
  db, err = gorm.Open("sqlite3", "./db/gorm.db")
  if err != nil {
    fmt.Println(err)
  }

  db.AutoMigrate(&Person{})

  // 连接池最小连接数
  //
  db.DB().SetMaxIdleConns(10)
  // 连接池最大连接数
  //
  db.DB().SetMaxOpenConns(100)

  persons := routerVersionGroup.Group("/persons")
  persons.GET("/", getPersons)
  persons.GET("/:id", getPerson)
  persons.POST("/", createPerson)
  persons.PUT("/:id", updatePerson)
}

func updatePerson(c *gin.Context) {
  var person Person
  id := c.Params.ByName("id")
  if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
  }
  c.BindJSON(&person)
  db.Save(&person)
  c.JSON(200, person)
}

func createPerson(c *gin.Context) {
  var person Person
  c.BindJSON(&person)
  db.Create(&person)
  c.JSON(200, person)
}

func getPerson(c *gin.Context) {
  id := c.Params.ByName("id")
  var person Person
  if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
  } else {
    c.JSON(200, person)
  }
}

func getPersons(c *gin.Context) {
  var people []Person
  if err := db.Find(&people).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
  } else {
    c.JSON(200, people)
  }
}