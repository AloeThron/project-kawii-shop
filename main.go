package main

import (
	"os"

	"github.com/AloeThron/project-kawii-shop/config"
	"github.com/AloeThron/project-kawii-shop/modules/servers"
	"github.com/AloeThron/project-kawii-shop/packages/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
