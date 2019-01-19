package functions

import (
	"crypto/sha1"
	"encoding/hex"
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
	var status = ""
	if responseCode == "00" {
		status = "success"
	} else {
		status = "failed"
	}
	CreateLog(status)
	return result
}
func CreateHash(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
