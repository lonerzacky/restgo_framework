package routers

import (
	"../controllers"
	"../functions"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, functions.GetResponse("00", "Service RestFul API GoLang Version "+os.Getenv("APP_VERSION")))
	})

	r.GET("/listRole", controllers.ListRole)
	r.POST("/insertRole", controllers.InsertRole)
	r.PUT("/updateRole/:sysrole_kode", controllers.UpdateRole)
	r.DELETE("/deleteRole/:sysrole_kode", controllers.DeleteRole)

	r.GET("/listUser", controllers.ListUser)
	r.POST("/insertUser", controllers.InsertUser)
	r.PUT("/updateUser/:sysuser_id", controllers.UpdateUser)
	r.DELETE("/deleteUser/:sysuser_id", controllers.DeleteUser)

	return r
}
