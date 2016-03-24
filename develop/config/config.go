package config

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File = nil

func init() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("not found config.ini!")
	}

}
func GetConfigKey(section, key string) *ini.Key {
	val, err := cfg.Section(section).GetKey(key)
	if err != nil {
		return nil
	}
	return val
}
func DB(key string) *ini.Key {
	return GetConfigKey("db", key)
}
