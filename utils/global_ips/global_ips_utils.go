package global_ips

import (
	"TCP_Packet/models"
	"sync"
)

var SharedIPS = make(map[string]struct{})
var lock sync.RWMutex

func GetIPS() map[string]struct{} {
	lock.RLock()
	defer lock.RUnlock()
	return SharedIPS
}

func UpdateIPS(ips *models.IPS) {
	lock.Lock()
	defer lock.Unlock()
	//sharedVariable++
	for _, ip := range ips.ListIp {
		SharedIPS[ip] = struct{}{}
	}
}
func DeleteIPS(ips *models.IPS) {
	lock.Lock()
	defer lock.Unlock()
	//sharedVariable++
	for _, ip := range ips.ListIp {
		delete(SharedIPS, ip)
	}
}
