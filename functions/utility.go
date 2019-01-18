package functions

import (
	"crypto/sha1"
	"encoding/hex"
	"net"
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

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
