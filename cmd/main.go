package main

import (
	"os"

	"github.com/AyBalatsan/TimeTracker/configs"
	"github.com/AyBalatsan/TimeTracker/pkg/handler"
	"github.com/AyBalatsan/TimeTracker/pkg/repository"
	"github.com/AyBalatsan/TimeTracker/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// инициализируем зависимости
	configs.ConfigurationInstaller()

	// подключение к базе данных Postgres
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		DBName:   viper.GetString("database.dbname"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("database.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error init DB %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(Server)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		logrus.Printf("error occured while running http server: %s", err.Error())
	}
}
