package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"tech-ex-server/internal/app/apiserver"
	"tech-ex-server/internal/app/model"
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

	stConfig := storage.NewConfig()
	sf := storage.New(stConfig)
	sf.CreateSchema()
	go func() {
		duration := time.Second * time.Duration(30)
		tk := time.NewTicker(duration)
		for range tk.C {
			rJson, _ := json.Marshal(apiserver.EditedData())

			it := model.Item{
				CreateAt: time.Now(),
				JsonData: rJson,
			}
			sf.Item().DropItem()
			er := sf.Item().CreateItem(it)

			fmt.Println(er)

			sf.Close()
		}
	}()
	s := apiserver.New(config)
	if err := s.StartServer(); err != nil {
		log.Fatal(err)
	}
}
