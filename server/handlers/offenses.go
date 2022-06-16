package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/db/repositories"
	"github.com/Heroin-lab/taxi_service.git/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) getAllUserOffenses(c *gin.Context) {
	userId := c.Query("user_id")

	offenses, err := r.services.GetAllUserOffenses(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.OffensesResponse{
		Data: offenses,
	})
}

func (r *Router) createUserOffense(c *gin.Context) {
	var inputModel repositories.Offense

	if err := c.BindJSON(&inputModel); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offenseId, err := r.services.CreateUserOffense(inputModel)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: offenseId,
	})
}

func (r *Router) updateUserOffense(c *gin.Context) {
	var inputModel repositories.Offense

	if err := c.BindJSON(&inputModel); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offenseId, err := r.services.UpdateUserOffense(inputModel)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: offenseId,
	})
}

func (r *Router) deleteUserOffenses(c *gin.Context) {
	userId := c.Query("offense_id")

	offenseId, err := r.services.DeleteUserOffense(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: offenseId,
	})
}
