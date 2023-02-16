package file

import (
	"TCP_Packet/utils/global_ips"
	"TCP_Packet/utils/settings"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
	"time"
)

func ConvertSetToString(setContent map[string]struct{}) []string {
	var ips []string
	for ip, _ := range setContent {
		ips = append(ips, ip)
	}
	return ips
}
func WriteIPSToFile(setContent map[string]struct{}) {
	ips := ConvertSetToString(setContent)
	ipsStr := strings.Join(ips, ",")
	err := ioutil.WriteFile(settings.GetIPConfig().NameFile, []byte(ipsStr), 0644)
	if err != nil {
		glog.Error("Error write ips to file: ", err)
	}
}

func ReadIPSFromFile() map[string]struct{} {
	var setIPs = make(map[string]struct{})
	var ips []string
	data, err := ioutil.ReadFile(settings.GetIPConfig().NameFile)
	if err != nil {
		fmt.Println("Error read IPS from file:", err)
	} else {
		ips = strings.Split(string(data), ",")
		for _, ip := range ips {
			setIPs[ip] = struct{}{}
		}
	}
	return setIPs
}
func ScheduleWriteIPSToFile() {
	glog.Info("Scheduled Write IPS to File at Midnight Daily")
	ticker := time.NewTicker(24 * time.Hour) // create a ticker that ticks every 24 hours
	defer ticker.Stop()

	for {
		now := time.Now()                                                                     // get the current time
		midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // calculate the next midnight
		if now.Before(midnight) {
			<-time.After(midnight.Sub(now)) // wait until midnight
		}
		WriteIPSToFile(global_ips.GetIPS())
	}
}
