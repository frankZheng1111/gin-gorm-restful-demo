package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(routerVersionGroup *gin.RouterGroup) {
	pages := routerVersionGroup.Group("/pages")
	initPagesRouters(pages)
}
