package functions

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetResponse(response_code string, response_message string) interface{} {
	type Response struct {
		Response_code    string `json:"response_code"`
		Response_message string `json:"response_message"`
	}

	result := Response{
		response_code,
		response_message,
	}
	return result
}

func GetResponseWithData(response_code string, response_message string, response_data interface{}) interface{} {
	type Response struct {
		Response_code    string      `json:"response_code"`
		Response_message string      `json:"response_message"`
		Response_data    interface{} `json:"response_data"`
	}

	result := Response{
		response_code,
		response_message,
		response_data,
	}
	return result
}
func CreateHash(key string) string {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
