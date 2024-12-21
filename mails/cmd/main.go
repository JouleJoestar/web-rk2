package main

import (
	"flag"
	"log"
	"mail/internal/api"
	"mail/internal/config"
	"mail/internal/provider"
	"mail/internal/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "../configs/example.yaml", "путь к файлу конфигурации")
	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	uc := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, uc, cfg)
	srv.Run()

}
