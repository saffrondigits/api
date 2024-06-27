package controller

import (
	"net/http"

	"github.com/saffrondigits/api/models"
	"github.com/saffrondigits/api/repo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiController struct {
	dbConn repo.SqlDbImplementation
	logger *logrus.Logger
}

func NewApiController(dbConn repo.SqlDbImplementation, logger *logrus.Logger) *ApiController {
	return &ApiController{
		dbConn: dbConn,
		logger: logger,
	}
}

func SetUpRoute(router *gin.Engine, apiCtrl *ApiController) *gin.Engine {

	router.GET("/ping", apiCtrl.Ping)
	router.POST("/register", apiCtrl.Register)

	return router
}
func (ac *ApiController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ac *ApiController) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		ac.logger.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.dbConn.RegisterUser(user)
	if err != nil {
		ac.logger.Error("Error registering user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
