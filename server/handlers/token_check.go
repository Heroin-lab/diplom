package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (r *Router) CheckToken(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusUnauthorized, "Empty header")
		return
	}

	headerParse := strings.Split(header, " ")
	if len(headerParse) != 2 {
		response.NewErrorResponse(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := r.services.ParseToken(headerParse[1])
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
