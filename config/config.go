package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

// グローバル参照とする
var Config ConfigList

// main関数より前に実行するように設定
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	//iniファイル読み込み
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	//読み込んだiniファイルを設定
	Config = ConfigList{
		Port:      cfg.Section("web").Key("post").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
