package controllers

import (
	"TCP_Packet/models"
	"TCP_Packet/services"
	"TCP_Packet/utils/http_utils"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"net/http"
)

func UpdateIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IPS)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&requestBody)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			status, res := http_utils.ResponseError("Json syntax error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		case *json.UnmarshalTypeError:
			status, res := http_utils.ResponseError("Json Unmarshal Type Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		default:
			status, res := http_utils.ResponseError("Unknown Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		}
	} else {
		if len(requestBody.ListIp) < 1 {
			err = errors.New("Length of IPS is zero, try again")
			status, res := http_utils.ResponseError("Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		} else {
			status, res := services.UpdateIPS(requestBody)
			w.WriteHeader(status)
			w.Write(res)
		}
	}
}
func DeleteIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IPS)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&requestBody)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			status, res := http_utils.ResponseError("Json syntax error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		case *json.UnmarshalTypeError:
			status, res := http_utils.ResponseError("Json Unmarshal Type Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		default:
			status, res := http_utils.ResponseError("Unknown Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		}
	} else {
		if len(requestBody.ListIp) < 1 {
			err = errors.New("Length of IPS is zero, try again")
			status, res := http_utils.ResponseError("Error", err, http.StatusBadRequest)
			w.WriteHeader(status)
			w.Write(res)
		} else {
			status, res := services.DeleteIPS(requestBody)
			w.WriteHeader(status)
			w.Write(res)
		}
	}
}
func GetIPS(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	status, res := services.GetIPS()
	w.WriteHeader(status)
	w.Write(res)
}
