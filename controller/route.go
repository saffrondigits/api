package controller

import (
	"net/http"

	"github.com/saffrondigits/api/models"
	"github.com/saffrondigits/api/repo"
	"github.com/saffrondigits/api/security"

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
	router.POST("/login", apiCtrl.Login)
	router.DELETE("/delete", apiCtrl.DeleteUser)

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

	// Encrypt password to hash
	hash, err := security.GenerateHash(user.Password)
	if err != nil {
		ac.logger.Error("Error generating password hash: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hash

	err = ac.dbConn.RegisterUser(user)
	if err != nil {
		ac.logger.Error("Error registering user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (ac *ApiController) Login(c *gin.Context) {
	var loginCred models.LoginCred

	if err := c.ShouldBindJSON(&loginCred); err != nil {
		ac.logger.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email exists
	dbCred, err := ac.dbConn.CheckIfEmailExists(loginCred.Email)
	if err != nil {
		ac.logger.Error("error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if password matches
	if !security.CompareHashAndPassword(dbCred.Hash, loginCred.Password) {
		ac.logger.Error("error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username/password wrong"})
		return
	}

	// Acknowledge user they are logged in
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

func (ac *ApiController) DeleteUser(c *gin.Context) {
	var loginCred models.LoginCred

	if err := c.ShouldBindJSON(&loginCred); err != nil {
		ac.logger.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email exists
	_, err := ac.dbConn.CheckIfEmailExists(loginCred.Email)
	if err != nil {
		ac.logger.Error("error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//delete the user
	err = ac.dbConn.DeleteCheckedUser(loginCred.Email)
	if err != nil {
		ac.logger.Error("Error Deleting user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
