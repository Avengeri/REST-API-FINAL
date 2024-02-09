package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"start/internal/service"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	headerPorts := strings.Split(header, " ")
	if len(headerPorts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}
	userId, err := service.ParseToken(headerPorts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}