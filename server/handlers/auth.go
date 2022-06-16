package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/Heroin-lab/taxi_service.git/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Router) signUp(c *gin.Context) {
	var input request.User

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: id,
	})
}

func (h *Router) signIn(c *gin.Context) {
	var input request.UserSignIn

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, _, err := h.services.Authorization.GenerateToken(input.PhoneNumber, input.Password, 10)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, id, err := h.services.Authorization.GenerateToken(input.PhoneNumber, input.Password, 60)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.TokenResponse{
		Id:           id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
