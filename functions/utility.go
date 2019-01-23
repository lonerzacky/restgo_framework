package functions

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetResponse(responseCode string, responseMessage string) interface{} {
	type Response struct {
		ResponseCode    string `json:"response_code"`
		ResponseMessage string `json:"response_message"`
	}

	result := Response{
		responseCode,
		responseMessage,
	}
	return result
}

func GetResponseWithData(responseCode string, responseMessage string, responseData interface{}) interface{} {
	type Response struct {
		ResponseCode    string      `json:"response_code"`
		ResponseMessage string      `json:"response_message"`
		ResponseData    interface{} `json:"response_data"`
	}

	result := Response{
		responseCode,
		responseMessage,
		responseData,
	}
	return result
}

func GetResponseWithDataLogging(context *gin.Context, responseCode string, responseMessage string, responseData interface{}) interface{} {
	type Response struct {
		ResponseCode    string      `json:"response_code"`
		ResponseMessage string      `json:"response_message"`
		ResponseData    interface{} `json:"response_data"`
	}

	result := Response{
		responseCode,
		responseMessage,
		responseData,
	}
	RequestUrl := context.Request.Host
	RequestApi := context.Request.URL.String()
	Method := context.Request.Method
	Params := context.Request.Form
	ResponseApi, _ := json.Marshal(result)
	ParamsJSON, _ := json.Marshal(Params)
	Status := ""
	if responseCode == "00" {
		Status = "success"
	} else {
		Status = "failed"
	}
	CreateLog(RequestUrl, RequestApi, Method, ParamsJSON, ResponseApi, Status)
	return result
}

func CreateHash(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
