package main

import (
	"context"
	"os"
	"os/signal"
	"ozon/pkg/handler"
	"ozon/pkg/repository"
	"ozon/pkg/service"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	repos, err := selectDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
		return
	}

	services := service.NewService(repos)
	srv := handler.NewServer(viper.GetString("port"), services)

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("LinkApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

const PostreSQLDB = "postgres"
const MemoryDB = "memory"

func selectDB() (*repository.Repository, error) {
	if viper.GetString("paramdb") == PostreSQLDB {
		db, err := repository.NewPostgresDB(repository.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBSchema: viper.GetString("db.schema"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: viper.GetString("db.password"),
		})
		if err != nil {
			logrus.Fatalf("failed to initialize db: %s", err.Error())
			return nil, err
		}
		if err := db.Close(); err != nil {
			logrus.Errorf("error occured on db connection close: %s", err.Error())
		}
		return repository.NewRepository(db), nil
	}
	if viper.GetString("paramdb") == MemoryDB {
		return repository.NewMemoryRepository(), nil

	}
	return nil, errors.New("failed to params db")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
