package global_ips

import (
	"TCP_Packet/models"
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

//var (
//	SharedIPS = make(map[string]struct{})
//	mutex     sync.RWMutex
//)

//var lock sync.RWMutex
//Key is IP
func CheckIPSExsist(db *leveldb.DB, key []byte) bool {
	has, err := db.Has(key, nil)
	if err != nil {
		log.Fatal(err)
	}
	if has {
		return true
	} else {
		return false
	}
}
func UpdateIPSToLevelDB(db *leveldb.DB, ips *models.IPS) {
	for _, ip := range ips.ListIp {
		err := db.Put([]byte(ip), []byte("1"), nil)
		if err != nil {
			glog.Error(ip, err)
		}
	}
}
func DeleteIPSToLevelDB(db *leveldb.DB, ips *models.IPS) {
	for _, ip := range ips.ListIp {
		db.Delete([]byte(ip), nil)
	}
}
