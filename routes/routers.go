package routes

import (
	"example-api/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRoutes() *gin.Engine {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic("Error loading .env file")
	}

	path := os.Getenv("CONTEXT_PATH")

	r := gin.Default()
	// setup routes
	r.POST(path + "/role", controller.CreateRole)
	r.GET(path + "/get-roles", controller.GetRoles)
	r.GET(path + "/get-role-by-id", controller.GetRoleById)
	r.PUT(path + "/update-role", controller.UpdateRole)
	r.DELETE(path + "/delete-role", controller.DeleteRole)
	r.GET(path + "/get-roles-page", controller.GetRolePage)

	return r
}
