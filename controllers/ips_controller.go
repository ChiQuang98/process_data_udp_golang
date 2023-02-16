package controllers

import (
	"TCP_Packet/models"
	"TCP_Packet/services"
	"encoding/json"
	"net/http"
)

func UpdateIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IPS)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		switch err.(type) {
		case *json.SyntaxError:
			w.Write([]byte("Json syntax error:" + err.Error()))
		case *json.UnmarshalTypeError:
			w.Write([]byte("Json Unmarshal Type Error:" + err.Error()))
		default:
			w.Write([]byte("Unknown Error:" + err.Error()))
		}
	} else {
		//if it's a root server -> make call to sub Server on port 7002
		//if settings.GetRestful().Port != settings.GetRestful().SubPort {
		//	subURL := "update/v1/ips"
		//	host := settings.GetRestful().Host
		//	port := settings.GetRestful().SubPort
		//	http_utils.RequestUpdateIPS(host, port, subURL, http.MethodPost, requestBody)
		//}
		status, res := services.UpdateIPS(requestBody)
		w.WriteHeader(status)
		w.Write(res)
	}
}
func DeleteIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IPS)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	//for ip, _ := range global_ips.GetIPS() {
	//	print(ip + ",")
	//}
	//println("")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		switch err.(type) {
		case *json.SyntaxError:
			w.Write([]byte("Json syntax error:" + err.Error()))
		case *json.UnmarshalTypeError:
			w.Write([]byte("Json Unmarshal Type Error:" + err.Error()))
		default:
			w.Write([]byte("Unknown Error:" + err.Error()))
		}
	} else {
		//if settings.GetRestful().Port != settings.GetRestful().SubPort {
		//	subURL := "delete/v1/ips"
		//	host := settings.GetRestful().Host
		//	port := settings.GetRestful().SubPort
		//	http_utils.RequestUpdateIPS(host, port, subURL, http.MethodDelete, requestBody)
		//}
		status, res := services.DeleteIPS(requestBody)
		w.WriteHeader(status)
		w.Write(res)
	}
}
