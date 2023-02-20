package services

import (
	"TCP_Packet/models"
	"TCP_Packet/utils/global_ips"
	"TCP_Packet/utils/http_utils"
	"TCP_Packet/utils/leveldb_utils"
	"net/http"
)

func UpdateIPS(ips *models.IPS) (int, []byte) {
	levelDB := leveldb_utils.GetLevelDBConnection()
	err := global_ips.UpdateIPSToLevelDB(levelDB, ips)
	if err != nil {
		code, resJson := http_utils.ResponseError("ERROR", err, http.StatusInternalServerError)
		return code, resJson
	}
	code, resJson := http_utils.ResponseOK("SUCCESS: Updated IPS", http.StatusOK)
	return code, resJson
}
func DeleteIPS(ips *models.IPS) (int, []byte) {
	levelDB := leveldb_utils.GetLevelDBConnection()
	err := global_ips.DeleteIPSToLevelDB(levelDB, ips)
	if err != nil {
		code, resJson := http_utils.ResponseError("ERROR", err, http.StatusInternalServerError)
		return code, resJson
	}
	code, resJson := http_utils.ResponseOK("SUCCESS: Deleted IPS", http.StatusOK)
	return code, resJson
}
func GetIPS() (int, []byte) {
	levelDB := leveldb_utils.GetLevelDBConnection()
	err, ips := global_ips.ReadIPSFromLevelDB(levelDB)
	if err != nil {
		code, resJson := http_utils.ResponseError("ERROR", err, http.StatusInternalServerError)
		return code, resJson
	}
	code, resJson := http_utils.ResponseGetAllIPS(ips, http.StatusOK)
	return code, resJson
}
