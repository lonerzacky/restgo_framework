package functions

import (
	"github.com/vjeantet/jodaTime"
	"restgo_framework/database"
	"time"
)

var db = database.ConnectDB()

func CreateLog(RequestUrl string, RequestApi string, Method string, Params interface{}, Response interface{}, Status string) {
	tx := db.Begin()
	RequestTime := jodaTime.Format("YYYY-MM-dd hh:mm:ss", time.Now())
	tx.Exec("INSERT INTO logservice(request_url,request_api,method,params,request_time,response,status) VALUES (?,?,?,?,?,?,?)",
		RequestUrl, RequestApi, Method, Params, RequestTime, Response, Status)
	tx.Commit()
}
