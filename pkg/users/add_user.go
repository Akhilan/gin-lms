package users

import (
	"net/http"

	"github.com/akhilan/gin-lms/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	Name  string `json:"name"`
	Place string `json:"place"`
	Mail  string `json:"mail"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Name = body.Name
	user.Place = body.Place
	user.Mail = body.Mail

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}
