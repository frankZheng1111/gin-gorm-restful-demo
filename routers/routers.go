package routers

import (
	"gin-gorm-restful-demo/routers/groups"
	"github.com/gin-gonic/gin"
)

func InitRouters(routerVersionGroup *gin.RouterGroup) {
	persons := routerVersionGroup.Group("/persons")
	groups.InitPersonsRouters(persons)
}
