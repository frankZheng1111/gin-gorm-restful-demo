package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(routerVersionGroup *gin.RouterGroup) {
	persons := routerVersionGroup.Group("/persons")
	initPersonsRouters(persons)
}
