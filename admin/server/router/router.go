package router

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/router/api"
	"admin/server/router/system"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(mode string) (g *gin.Engine, err error) {
	err = app.SetupValidate()
	if err != nil {
		return nil, err
	}

	r := gin.New()
	gin.SetMode(mode)
	r.Use(ginzap.Ginzap(log.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log.Logger, true))

	authMiddleware, err := system.GetJwtMiddleWare(system.Login, system.Logout)
	if err != nil {
		return nil, err
	}
	r.NoRoute(authMiddleware.MiddlewareFunc(), system.NoRoute)

	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := authMiddleware.MiddlewareFunc
	systemApi := r.Group("/system", auth())
	{
		systemApi.GET("/user/refresh_token", authMiddleware.RefreshHandler)
		systemApi.POST("/user/logout", authMiddleware.LogoutHandler)
		systemApi.GET("/user/info", system.GetUserInfo)

		systemApi.POST("/user", system.AddUser)
		systemApi.GET("/user", system.GetUsers)
		systemApi.GET("/user/:id", system.GetUser)
		systemApi.PUT("/user/:id", system.UpdateUser)
		systemApi.PUT("/user", system.UpdateUserInfo)
		systemApi.PUT("/user/password/:id", system.UpdateUserPassword)
		systemApi.DELETE("/user/:id", system.DeleteUser)

		systemApi.GET("/email", system.GetEmail)
		systemApi.POST("/email", system.AddEmail)
		systemApi.PUT("/email/:id", system.UpdateEmail)

		systemApi.GET("/ldap", system.GetLdap)
		systemApi.POST("/ldap", system.AddLdap)
		systemApi.PUT("/ldap/:id", system.UpdateLdap)

		systemApi.GET("/txsms", system.GetTxsms)
		systemApi.POST("/txsms", system.AddTxsms)
		systemApi.PUT("/txsms/:id", system.UpdateTxsms)

		systemApi.POST("/msg", system.AddMsg)
		systemApi.GET("/msg", system.GetMsgs)
		systemApi.GET("/msg/:id", system.GetMsg)
		systemApi.PUT("/msg/:id", system.UpdateMsg)
		systemApi.DELETE("/msg/:id", system.DeleteMsg)
		systemApi.POST("/msg/:id/user", system.SendMsg)

	}

	apis := r.Group("/api", auth())
	{
		apis.POST("/db", api.AddDB)
		apis.GET("/db", api.GetDBs)
		apis.GET("/db/:id", api.GetDB)
		apis.PUT("/db/:id", api.UpdateDB)
		apis.DELETE("/db/:id", api.DeleteDB)

		apis.POST("/dbuser", api.AddDBUser)
		apis.GET("/dbuser", api.GetDBUsers)
		apis.GET("/dbuser/:id", api.GetDBUser)
		apis.PUT("/dbuser/:id", api.UpdateDBUser)
		apis.DELETE("/dbuser/:id", api.DeleteDBUser)
		apis.PUT("/dbuser/:id/db", api.UpdateDBUserDB)

		apis.POST("/ip", api.AddIP)
		apis.GET("/ip", api.GetIPs)
		apis.GET("/ip/:id", api.GetIP)
		apis.PUT("/ip/:id", api.UpdateIP)
		apis.DELETE("/ip/:id", api.DeleteIP)
	}

	return r, nil
}
