package http_utils

import (
	"TCP_Packet/models"
	"encoding/json"
	//"strconv"
)

func ResponseError(msg string, err error, code int) (int, []byte) {
	resStr := msg + ": " + err.Error()
	res := models.Response{Message: resStr}
	resJson, _ := json.Marshal(res)
	return code, resJson
}
func ResponseOK(msg string, code int) (int, []byte) {
	res := models.Response{Message: msg}
	resJson, _ := json.Marshal(res)
	return code, resJson
}
func ResponseGetAllIPS(ips models.IPS, code int) (int, []byte) {
	res := models.IPS{ListIp: ips.ListIp}
	resJson, _ := json.Marshal(res)
	return code, resJson
}
