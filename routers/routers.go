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

	r.GET("/listRole", controllers.AuthJWT, controllers.ListRole)
	r.POST("/insertRole", controllers.AuthJWT, controllers.InsertRole)
	r.PUT("/updateRole/:sysrole_kode", controllers.AuthJWT, controllers.UpdateRole)
	r.DELETE("/deleteRole/:sysrole_kode", controllers.AuthJWT, controllers.DeleteRole)

	r.GET("/listModul", controllers.AuthJWT, controllers.ListModul)
	r.POST("/insertModul", controllers.AuthJWT, controllers.InsertModul)
	r.PUT("/updateModul/:sysmodul_kode", controllers.AuthJWT, controllers.UpdateModul)
	r.DELETE("/deleteModul/:sysmodul_kode", controllers.AuthJWT, controllers.DeleteModul)

	r.POST("/insertRmodul", controllers.AuthJWT, controllers.InsertRmodul)
	r.DELETE("/deleteRmodul/:kodeRole/:kodeModul", controllers.AuthJWT, controllers.DeleteRmodul)

	r.GET("/listUser", controllers.AuthJWT, controllers.ListUser)
	r.POST("/insertUser", controllers.AuthJWT, controllers.InsertUser)
	r.PUT("/updateUser/:sysuser_id", controllers.AuthJWT, controllers.UpdateUser)
	r.DELETE("/deleteUser/:sysuser_id", controllers.AuthJWT, controllers.DeleteUser)

	r.POST("/verifyLogin", controllers.AuthJWT, controllers.VerifyLogin)
	r.POST("/changePassword", controllers.AuthJWT, controllers.ChangePassword)
	r.GET("/createToken", controllers.CreateTokenEndpoint)

	return r
}
