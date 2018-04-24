package main

import (
  "gin-gorm-restful-demo/routers"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  v1 := router.Group("/v1")
  routers.InitRouters(v1)

  router.Run() // 8080
}
