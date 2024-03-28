package routes

import (
	"github.com/fvnis/golang-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitiRoutes(r *gin.RouterGroup) {
  r.GET("/getUserById/:userId", controller.GetUserById)
  r.GET("/getUserByEmail/:userEmail", controller.GetUserByEmail)
  r.POST("/createUser", controller.CreateUser)
  r.PUT("/updateUser/:userId", controller.UpdateUser)
  r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
