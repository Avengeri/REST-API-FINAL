package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"start/internal/model"
	"strconv"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func (h *Handler) pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON data
// @Tags user
// @Param user body model.UserTodo true "User data in JSON format"
// @SecurityDefinitions.apikey
// @Security ApiKeyAuth
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Router /user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user *model.UserTodo

	err := c.ShouldBindJSON(&user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Todo.SetUserService(user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "the user has been successfully created"})
}

// GetUserByID godoc
// @Summary GetById a user
// @Description GetById a user with the provided JSON data
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @SecurityDefinitions.apikey
// @Security ApiKeyAuth
// @Success 200 {string} string "User get successfully"
// @Failure 500 {object} errorResponse
// @Router /user/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	userStrID := c.Param("id")
	userID, err := strconv.Atoi(userStrID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Todo.GetUserByIDService(userID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

// CheckUserByID godoc
// @Summary CheckById if a user exists
// @Description CheckById if a user with the provided ID exists
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @SecurityDefinitions.apikey
// @Security ApiKeyAuth
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /user/check/{id} [get]
func (h *Handler) CheckUserByID(c *gin.Context) {
	userStrID := c.Param("id")
	userID, err := strconv.Atoi(userStrID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	exist, err := h.service.Todo.CheckUserByIDService(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if exist {
		c.JSON(http.StatusOK, statusResponse{Status: "the user was successfully found"})
	} else {
		c.JSON(http.StatusNotFound, statusResponse{Status: "the user was not found"})
	}
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Delete a user with the provided ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @SecurityDefinitions.apikey
// @Security ApiKeyAuth
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Router /user/{id} [delete]
func (h *Handler) DeleteUserByID(c *gin.Context) {
	userStrID := c.Param("id")
	userID, err := strconv.Atoi(userStrID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Todo.DeleteUserByIdService(userID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "the user has been successfully deleted"})
}

// GetAllUsersIDs godoc
// @Summary GetById a list of all users
// @Description GetById a list of all users with their IDs
// @Tags user
// @Accept json
// @Produce json
// @SecurityDefinitions.apikey
// @Security ApiKeyAuth
// @Success 200 {array} int "List of user IDs"
// @Failure 400 {object} errorResponse
// @Router /user/get_all [get]
func (h *Handler) GetAllUsersIDs(c *gin.Context) {
	userIDs, err := h.service.Todo.GetAllUserIDService()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userIDs)
}
