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
}

func main() {
	cfg := Config{}
	loadenvconf.LoadEnvConfig("dotenv/.env", &cfg)
	fmt.Println(cfg)
}
