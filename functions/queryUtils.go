package functions

import (
	"../database"
	"github.com/gin-gonic/gin"
	"github.com/vjeantet/jodaTime"
	"time"
)

var db = database.ConnectDB()

func CreateLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin()
		tx.Exec("INSERT INTO logservice(request_url,request_api,method,params,request_time,response,status) VALUES (?,?,?,?,?,?,?)",
			c.Request.Host, c.Request.URL.String(), c.Request.Method, "params", jodaTime.Format("YYYY-MM-dd hh:mm:ss", time.Now()), "response", "status")
		tx.Commit()
	}
}
