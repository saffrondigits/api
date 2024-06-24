package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/api/db"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Infof("Starting the server at address %s", Address)
	r.Run(Address)
}
