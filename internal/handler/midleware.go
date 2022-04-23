package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "auth not provided")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		h.log.Errorf("ParseToken err: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Sprintf("ParseToken err: %s", err.Error()))
		return
	}

	c.Set(userIdCtx, userId)

}
