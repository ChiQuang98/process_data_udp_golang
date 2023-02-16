package services

import (
	"TCP_Packet/models"
	"TCP_Packet/utils/global_ips"
	"TCP_Packet/utils/leveldb_utils"
	"net/http"
)

func UpdateIPS(ips *models.IPS) (int, []byte) {
	levelDB := leveldb_utils.GetLevelDBConnection()
	global_ips.UpdateIPSToLevelDB(levelDB, ips)
	res := []byte("Accepted")
	return http.StatusOK, res
}
func DeleteIPS(ips *models.IPS) (int, []byte) {
	levelDB := leveldb_utils.GetLevelDBConnection()
	global_ips.DeleteIPSToLevelDB(levelDB, ips)
	res := []byte("Accepted")
	return http.StatusOK, res
}
