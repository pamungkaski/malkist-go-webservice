package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pamungkaski/malkist-go"
	"net/http"
)

type DistanceMatrixRequest struct {
	Origins      []string `form:"origins"`
	Destinations []string `form:"destinations"`
	Key          string   `form:"key"`
}

func (s Service) CalculateDistance(ctx *gin.Context) {
	var req DistanceMatrixRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := malkist.Malkist{
		Key: s.Key,
	}

	result, err := m.CalculateDistance(req.Origins, req.Destinations)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
