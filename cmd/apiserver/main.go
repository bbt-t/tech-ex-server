package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"tech-ex-server/internal/app/apiserver"
	"tech-ex-server/internal/app/storage"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(
		&configPath,
		"configs-path",
		"configs/apiserver.toml",
		"config-file path",
	)
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("---> Looks like you don't have Postgres DataBase running <---", err)
		}
	}()
	apiserver.ExitHandler()

	stConfig := storage.NewConfig()
	sf := storage.New(stConfig)
	sf.CreateSchema()

	duration := time.Second * time.Duration(30)
	go apiserver.MakeItem(duration, sf)

	s := apiserver.New(config)
	if err := s.StartServer(); err != nil {
		log.Fatal(err)
	}
}
