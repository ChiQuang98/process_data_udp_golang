package global_ips

import (
	"TCP_Packet/models"
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

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
func UpdateIPSToLevelDB(db *leveldb.DB, ips *models.IPS) error {
	for _, ip := range ips.ListIp {
		err := db.Put([]byte(ip), []byte("1"), nil)
		if err != nil {
			glog.Error(ip, err)
			return err
		}
	}
	return nil
}
func ReadIPSFromLevelDB(db *leveldb.DB) (error, models.IPS) {
	// Create an iterator for the database
	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	ips := models.IPS{}
	var arrIPS []string
	// Iterate over all the key-value pairs in the database
	for iter.Next() {
		key := iter.Key()
		//value := iter.Value()
		arrIPS = append(arrIPS, string(key))
		//glog.Info("Key: %s, Value: %s\n", key, value)
	}

	if err := iter.Error(); err != nil {
		glog.Error("Error retrieving data from database:", err)
		return err, ips
	}
	ips.ListIp = arrIPS
	return nil, ips
}
func DeleteIPSToLevelDB(db *leveldb.DB, ips *models.IPS) error {
	for _, ip := range ips.ListIp {
		err := db.Delete([]byte(ip), nil)
		if err != nil {
			glog.Error(ip, err)
			return err
		}
	}
	return nil
}
