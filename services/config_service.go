package services

import (
	"TCP_Packet/models"
	"TCP_Packet/utils/global_ips"
	"net/http"
)

func UpdateIPS(ips *models.IPS) (int, []byte) {
	global_ips.UpdateIPS(ips)
	res := []byte("Accepted")
	return http.StatusOK, res
}
func DeleteIPS(ips *models.IPS) (int, []byte) {
	global_ips.DeleteIPS(ips)
	res := []byte("Accepted")
	return http.StatusOK, res
}
