package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sug5806/bookstore_users-api/domain/users"
	"github.com/sug5806/bookstore_users-api/services"
	"github.com/sug5806/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestError{
			Status:  http.StatusBadRequest,
			Message: "invalid json body",
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO:Handle error
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
