package http_utils

import (
	"TCP_Packet/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	//"strconv"
)

func RequestUpdateIPS(host string, port int, subURL string, method string, ips *models.IPS) {
	var serverBasePath string = fmt.Sprintf("http://%s:%d/%s", host, port, subURL)
	//urlStr := "http://" + host + ":" + strconv.Itoa(port) + subURL
	//escapedUrl, _ := url.Parse(urlStr)
	println(serverBasePath)
	//data := []byte(`{"key": "value"}`)
	data, err := json.Marshal(&ips)
	if err != nil {
		glog.Error(err)
	}
	// NOTE this !!

	glog.Info(string(data))
	req, err := http.NewRequest(method, serverBasePath, bytes.NewBuffer(data))
	req.Close = true
	if err != nil {
		glog.Error("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{},
	}
	resp, err := client.Do(req)
	if err != nil {
		glog.Error("Error 2:", err)
		return
	}
	defer resp.Body.Close()
	glog.Info(resp.Status)
}
