package leveldb_utils

import (
	"TCP_Packet/utils/settings"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)

//func GetLevelDBConnection() *leveldb.DB {
//	db, err := leveldb.OpenFile("/u01/app_mobileid/golang/TCP_Packet/dbips", nil)
//	if err != nil {
//		glog.Error("Err levelDB: ", err)
//		//panic(err)
//	}
//	//defer db.Close()
//	return db
//}

var db *leveldb.DB
var once sync.Once

// OpenDB opens a LevelDB database
func GetLevelDBConnection() *leveldb.DB {
	var err error
	once.Do(func() {
		db, err = leveldb.OpenFile(settings.GetIPConfig().LevelDB, nil)
	})
	return db
}
