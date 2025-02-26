package routes

import (
	"example-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	// setup routes
	r.POST("/role", controller.CreateRole)
	r.GET("/get-roles", controller.GetRoles)
	r.GET("/get-role-by-id", controller.GetRoleById)
	r.PUT("/update-role", controller.UpdateRole)
	r.DELETE("/delete-role", controller.DeleteRole)

	return r

}
