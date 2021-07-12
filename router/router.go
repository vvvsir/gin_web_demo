package router

import (
	"gin_web_demo/controller"
	"gin_web_demo/logger"
	"gin_web_demo/middlewares"
	"net/http"

	_ "gin_web_demo/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter(model string) *gin.Engine {
	if model == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/updateuser", controller.UpdateUser)
		v1.POST("/createuser", controller.CreateUser)
		v1.GET("/deleteuser/:userId", controller.DeleteUserByUserId)
		v1.POST("/login", controller.Login)
		v1.POST("/getusers", controller.GetUserList)

		v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
		{
			v1.GET("/getuser/:userId", controller.GetUserByUserId)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
