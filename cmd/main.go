package main

import (
	server "coronairis-lp_web-app"
	"coronairis-lp_web-app/configs"
	"coronairis-lp_web-app/pkg/handler"
	"coronairis-lp_web-app/pkg/repository"
	"coronairis-lp_web-app/pkg/service"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "main",
			"error":    err,
		}).Fatalf("error initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "initConfig",
			"error":    err,
		}).Fatalf("error loading env variables")
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "main",
			"error":    err,
		}).Fatalf("failed to initialize database")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if viper.GetBool("logging.logfile") { // set output to log files if param logfile is true
		logName := strconv.FormatInt(time.Now().Unix(), 10)
		logDir := "logs"
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			err := os.Mkdir(logDir, 0744)
			if err != nil {
				log.WithFields(log.Fields{
					"package": "main",
					"file":    "main.go",
					"func":    "main",
					"message": err,
				}).Errorf("failed to create 'logs' dir")
				viper.Set("logging.logfile", false) // if the 'logs' directory could not be created, sets the logfile parameter to false
			}
		}
		logFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", logDir, logName), os.O_CREATE|os.O_WRONLY, 0777)
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	}

	if strings.EqualFold(viper.GetString("logging.format"), "json") {
		log.SetFormatter(new(log.JSONFormatter))
	}

	app := new(server.Server)
	if err := app.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.WithFields(log.Fields{
			"package":  "main",
			"file":     "main.go",
			"function": "main",
			"error":    err,
		}).Fatalf("error occured while running http server")
	}
}
