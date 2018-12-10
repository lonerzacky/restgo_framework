package routers

import (
	"../controllers"
	"../functions"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, functions.GetResponse("00", "Service RestFul API GoLang Version "+os.Getenv("APP_VERSION")))
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/listRole", controllers.ListRole)
	r.POST("/insertRole", controllers.InsertRole)
	r.PUT("/updateRole/:sysrole_kode", controllers.UpdateRole)
	r.DELETE("/deleteRole/:sysrole_kode", controllers.DeleteRole)

	r.GET("/listUser", controllers.ListUser)
	r.POST("/insertUser", controllers.InsertUser)
	r.PUT("/updateUser/:sysuser_id", controllers.UpdateUser)
	r.DELETE("/deleteUser/:sysuser_id", controllers.DeleteUser)

	r.POST("/verifyLogin", controllers.VerifyLogin)
	r.POST("/changePassword", controllers.ChangePassword)

	return r
}
