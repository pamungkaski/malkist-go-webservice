package service

import "github.com/gin-gonic/gin"

type Service struct {
	Router *gin.Engine
	Key    string
}

func NewService(key string) *Service {
	router := gin.Default()
	return &Service{
		Router: router,
		Key:    key,
	}
}

func (s Service) RegisterService() {
	s.Router.GET("/api/distance-matrix", func(ctx *gin.Context) { s.CalculateDistance(ctx) })
}

func (s Service) StartService(port string) {
	s.Router.Run(port)
}
