package controllers

import (
	"TCP_Packet/models"
	"TCP_Packet/services"
	"TCP_Packet/utils/global_ips"
	"encoding/json"
	"fmt"
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
		status, res := services.UpdateIPS(requestBody)
		w.WriteHeader(status)
		w.Write(res)
	}
}
func DeleteIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IPS)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	for i, number := range global_ips.GetIPS() {
		fmt.Printf("Index: %d, Value: %d\n", i, number)
	}
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
		status, res := services.DeleteIPS(requestBody)
		w.WriteHeader(status)
		w.Write(res)
	}
}
