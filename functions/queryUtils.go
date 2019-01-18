package functions

import (
	"../database"
	"github.com/gin-gonic/gin"
	"github.com/vjeantet/jodaTime"
	"time"
)

var db = database.ConnectDB()

func CreateLog(status string) (context *gin.Context) {
	tx := db.Begin()
	tx.Exec("INSERT INTO logservice(uri,method,params,ip_address,request_time,response,status) VALUES (?,?,?,?,?,?,?)",
		"url", "aa", "aa", "aa", jodaTime.Format("YYYY-MM-dd hh:mm:ss", time.Now()), "aa", status)
	tx.Commit()
	return
}
