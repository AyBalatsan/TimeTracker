package handler

import (
	"net/http"

	"github.com/AyBalatsan/TimeTracker/models"
	"github.com/gin-gonic/gin"
)

type getAllUsersResponse struct {
	Data []models.People `json:"data"`
}

func (h *Handler) createUser(c *gin.Context) {
	var input models.People
	// парсит нашу структуру и где binding:"required" не пропустит если поле пустое
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Users.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUser()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}
func (h *Handler) getUserByPassport(c *gin.Context) {

}
func (h *Handler) getUserTasksList(c *gin.Context) {

}
func (h *Handler) updateUser(c *gin.Context) {

}
func (h *Handler) deleteUser(c *gin.Context) {

}
