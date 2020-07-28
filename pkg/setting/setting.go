package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type Server struct {
	Host string
	HttpPort string
	RunMode string
	TimeZone string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	Username string
	Password string
	Database string
	Location string
}

var DatabaseSetting = &Database{}


var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}