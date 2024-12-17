package main

import (
	"flag"
	"log"
	"mail/internal/api"
	"mail/internal/config"
	"mail/internal/provider"
	"mail/internal/usecase" // Импортируем пакет usecase

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "../configs/example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем провайдер для работы с базой данных
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)

	// Создаем новый usecase
	uc := usecase.NewUsecase(prv) // Создаем новый usecase, передавая провайдер

	// Создаем новый сервер API
	srv := api.NewServer(cfg.IP, cfg.Port, uc, cfg) // Передаем usecase в сервер API

	// Запускаем сервер
	srv.Run()

}
