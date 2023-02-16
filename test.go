package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	db, err := leveldb.OpenFile("/u01/app_mobileid/golang/TCP_Packet/dbips", nil)
	if err != nil {
		glog.Error("Err levelDB: ", err)
		//panic(err)
	}
	key := []byte("mykey")
	value := []byte("myvalue")

	err = db.Put(key, value, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Stored key %s with value %s in LevelDB\n", string(key), string(value))
}
