package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/api/controller"
	"github.com/saffrondigits/api/db"
	"github.com/saffrondigits/api/repo"
	"github.com/sirupsen/logrus"
)

const (
	Address = "127.0.0.1:5000"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Infof("Establishing db connection...")

	dbInst, err := db.ConnectTOTheDatabase()
	if err != nil {
		log.Errorf("Error establishing db connection: %v", err)
		return
	}

	log.Infof("Successfully connected to the db!")

	_ = dbInst

	sqlImpl := repo.NewSqlDbImplementation(dbInst)
	apiCtrl := controller.NewApiController(sqlImpl, log)
	r := gin.Default()

	router := controller.SetUpRoute(r, apiCtrl)

	log.Infof("Starting the server at address %s", Address)
	router.Run(Address)
}
