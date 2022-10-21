package users

import (
	"net/http"

	"github.com/akhilan/gin-lms/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	Name  string `json:"name"`
	Place string `json:"place"`
	Mail  string `json:"mail"`
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := UpdateUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	user.Name = body.Name
	user.Place = body.Place
	user.Mail = body.Mail

	h.DB.Save(&user)

	c.JSON(http.StatusOK, &user)
}
