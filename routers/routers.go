package routers

import (
	"../controllers"
	"../functions"
	"../security"
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

	r.GET("/listRole", security.AuthJWT, controllers.ListRole)
	r.POST("/insertRole", security.AuthJWT, controllers.InsertRole)
	r.PUT("/updateRole/:sysrole_kode", security.AuthJWT, controllers.UpdateRole)
	r.DELETE("/deleteRole/:sysrole_kode", security.AuthJWT, controllers.DeleteRole)

	r.GET("/listModul", security.AuthJWT, controllers.ListModul)
	r.POST("/insertModul", security.AuthJWT, controllers.InsertModul)
	r.PUT("/updateModul/:sysmodul_kode", security.AuthJWT, controllers.UpdateModul)
	r.DELETE("/deleteModul/:sysmodul_kode", security.AuthJWT, controllers.DeleteModul)

	r.POST("/insertRmodul", security.AuthJWT, controllers.InsertRmodul)
	r.DELETE("/deleteRmodul/:kodeRole/:kodeModul", security.AuthJWT, controllers.DeleteRmodul)

	r.GET("/listUser", security.AuthJWT, controllers.ListUser)
	r.POST("/insertUser", security.AuthJWT, controllers.InsertUser)
	r.PUT("/updateUser/:sysuser_id", security.AuthJWT, controllers.UpdateUser)
	r.DELETE("/deleteUser/:sysuser_id", security.AuthJWT, controllers.DeleteUser)

	r.POST("/verifyLogin", security.AuthJWT, controllers.VerifyLogin)
	r.POST("/changePassword", security.AuthJWT, controllers.ChangePassword)
	r.GET("/createToken", security.CreateTokenEndpoint)

	return r
}
