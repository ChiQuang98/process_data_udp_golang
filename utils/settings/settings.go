package settings

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	GlogConfig   *GlogConfigs
	UDPCollector *UDPCollector
	UDPMapping   *UDPMapping
	Restful      *Restful
	System       *System
}
type GlogConfigs struct {
	LogDir  string
	MaxSize uint64
	V       int
}
type UDPCollector struct {
	Interface   string
	Port        uint64
	SrcUDP      string
	SizeCapture int32
}
type UDPMapping struct {
	Host string
	Port int
}
type Restful struct {
	Interface string
	Host      string
	Port      int
}
type System struct {
	NumberThread int
}

var settings Settings = Settings{}

func init() {
	//C:\Users\Admin\go\src\api_cv\mcu_api\setting.json
	//C:\Users\Admin\go\src\api_cv\lp-cms-client\setting.json
	//	content, err := ioutil.ReadFile("C:\\Users\\Admin\\go\\src\\api_cv\\mcu_api\\setting.json")
	//	content, err := ioutil.ReadFile("C:\\Users\\Admin\\go\\src\\api_cv\\lp-cms-client\\setting.json")
	//	abc:= filepath.Join(NamePackage)

	//absPath, _ := filepath.Abs(filepath.Join("mcu_api","setting.json"))
	//fmt.Println(abc)
	content, err := ioutil.ReadFile("setting.json")
	if err != nil {
		panic(err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		panic(jsonErr)
	}
}
func GetGlogConfig() *GlogConfigs {
	return settings.GlogConfig
}
func GetUDPCollector() *UDPCollector {
	return settings.UDPCollector
}
func GetUDPMapping() *UDPMapping {
	return settings.UDPMapping
}
func GetRestful() *Restful {
	return settings.Restful
}
func GetSystem() *System {
	return settings.System
}
