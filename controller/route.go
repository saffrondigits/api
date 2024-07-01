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
	router.POST("/login", apiCtrl.Login)

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
	if loginCred.Password != dbCred.Hash {
		ac.logger.Error("Error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username/password wrong"})
		return
	}

	// Acknowledge user they are logged in
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

func (as *ApiController) DeleteUser(c *gin.Context) {

}

// id | first_name | last_name |      email       | password
// ----+------------+-----------+------------------+----------
//   1 | John       | Mayer     | jm@example.com   | abcd
//   2 | Rahul      | Gandhi    | raga@example.com | abcd
