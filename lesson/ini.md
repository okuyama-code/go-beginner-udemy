package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

//ini

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {

	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{ //Configに値を入れていく
		Port: cfg.Section("web").Key("port").MustInt(8080),

		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {

	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)

}

// $ go run main.go
// Port = 8080
// DbName = webapp.sql
// SQLDriver = sqlite3
