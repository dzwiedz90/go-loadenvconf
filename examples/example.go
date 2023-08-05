package main

import (
	"fmt"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
)

type Config struct {
	SECRET_KEY  string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DUPA        string
}

func main() {
	cfg := Config{}
	_, err := loadenvconf.LoadEnvConfig("loadenvconf/.env", &cfg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}
}
