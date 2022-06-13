package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/server/request"
	"github.com/Heroin-lab/taxi_service.git/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) getAllCabMans(c *gin.Context) {
	cabMansList, err := r.services.CabMansServices.GetAllCabMans()
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.AllCabMansResponse{
		Data: cabMansList,
	})
}

func (r *Router) getOneCabMan(c *gin.Context) {
	cabManId := c.Param("id")

	cabManInfo, err := r.services.CabMansServices.GetOneCabMan(cabManId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.OneCabManResponse{
		Data: cabManInfo,
	})
}

func (r *Router) createCabMan(c *gin.Context) {
	var cabMan request.CabMan

	if err := c.BindJSON(&cabMan); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newCabMan, err := r.services.CabMansServices.CreateNewCabMan(cabMan)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: newCabMan,
	})
}

func (r *Router) updateCabMan(c *gin.Context) {
	var cabMan request.CabMan

	cabManId := c.Param("id")

	if err := c.BindJSON(&cabMan); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updId, err := r.services.CabMansServices.UpdateCabMan(cabManId, cabMan)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: updId,
	})
}

func (r *Router) deleteCabMan(c *gin.Context) {
	cabManId := c.Param("id")

	delId, err := r.services.CabMansServices.DeleteCabMan(cabManId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{
		Id: delId,
	})
}
