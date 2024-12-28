package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hpaes/go-api-final-project/src/api/controller"
)

type (
	GinRouter interface {
		SetAppHandlers()
		GetRouter() *gin.Engine
	}

	ginEngine struct {
		router                  *gin.Engine
		registerContoller       *controller.RegisterUserController
		getUserDetailController *controller.GetUserDetailController
		removeUserController    *controller.RemoveUserController
		listUsersController     *controller.ListUsersController
		updateUserController    *controller.UpdateUserController
	}
)

func NewGinEngine(router *gin.Engine,
	registerController *controller.RegisterUserController,
	getUserDetailController *controller.GetUserDetailController,
	removeUserController *controller.RemoveUserController,
	listUsersController *controller.ListUsersController,
	updateUserController *controller.UpdateUserController) *ginEngine {
	return &ginEngine{
		router:                  router,
		registerContoller:       registerController,
		getUserDetailController: getUserDetailController,
		removeUserController:    removeUserController,
		listUsersController:     listUsersController,
		updateUserController:    updateUserController,
	}
}

func (ge *ginEngine) SetAppHandlers() {
	ge.router.POST("/user", ge.registerUser())
	ge.router.GET("/user/:userId", ge.getUserDetail())
	ge.router.DELETE("/user/:userId", ge.removeUser())
	ge.router.GET("/users", ge.listUsers())
	ge.router.PUT("/user", ge.updateUser())
}

func (ge *ginEngine) updateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ge.updateUserController.Execute(c.Writer, c.Request)
	}
}

func (ge *ginEngine) listUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		query.Add("page", c.Param("page"))
		c.Request.URL.RawQuery = query.Encode()
		ge.listUsersController.Execute(c.Writer, c.Request)
	}
}

func (ge *ginEngine) removeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		query.Add("userId", c.Param("userId"))
		c.Request.URL.RawQuery = query.Encode()
		ge.removeUserController.Execute(c.Writer, c.Request)
	}
}

func (ge *ginEngine) getUserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		query.Add("userId", c.Param("userId"))
		c.Request.URL.RawQuery = query.Encode()
		ge.getUserDetailController.Execute(c.Writer, c.Request)
	}
}

func (ge *ginEngine) registerUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ge.registerContoller.Execute(c.Writer, c.Request)
	}
}

func (ge *ginEngine) GetRouter() *gin.Engine {
	return ge.router
}
