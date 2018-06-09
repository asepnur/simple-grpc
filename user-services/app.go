package main

import (
	"flag"
	"log"

	"github.com/asepnur/galahad/src/util/jsonconfig"
	"github.com/asepnur/simple-grpc/user-services/util/env"

	"github.com/asepnur/iskandar/src/util/conn"
	"github.com/asepnur/simple-grpc/user-services/webserver"
)

type configuration struct {
	Database  conn.DatabaseConfig `json:"database"`
	Redis     conn.RedisConfig    `json:"redis"`
	Webserver webserver.Config    `json:"webserver"`
}

// User struct :: save value
type User struct {
	UserID     int    `json:"id"`
	UserEmail  string `json:"email"`
	FullName   string `json:"name"`
	MSISDN     int    `json:"msisdn"`
	CreateTime string `json:"create_time"`
}

var (
	emptyTime  = "0001-01-01 00:00:00"
	layoutTime = "2006-01-02 15:04:05"
)

func main() {
	flag.Parse()

	//Load configuration
	cfgenv := env.Get()
	config := &configuration{}
	isLoaded := jsonconfig.Load(&config, "/etc", cfgenv) || jsonconfig.Load(&config, "./files/etc", cfgenv)
	if !isLoaded {
		log.Fatal("Failed to load configuration")
	}
	conn.InitDB(config.Database)
	webserver.Start(config.Webserver)
}
