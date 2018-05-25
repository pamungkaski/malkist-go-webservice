package service

import "github.com/gin-gonic/gin"

// GoogleMapsService is the service that will be wrapped and used in our service.
// This package will include all current and future functionality available.
type GoogleMapsService interface {
	CalculateDistance(ctx *gin.Context)
	CalculateElevation(ctx *gin.Context)
	CalculateGeolocation(ctx *gin.Context)
	CalculateGeocodingn(ctx *gin.Context)
	GetTimezone(ctx *gin.Context)
	GetDirections(ctx *gin.Context)
	GetPlaces(ctx *gin.Context)
	GetRoads(ctx *gin.Context)
}

// Service is the structure of malkist-go-webservice.
type Service struct {
	Router *gin.Engine
	Key    string
}

// New service will create a default gin router and include the API key to the struct.
func NewService(key string) *Service {
	router := gin.Default()
	return &Service{
		Router: router,
		Key:    key,
	}
}

// RegisterService will register each of every endpoint for the app.
func (s Service) RegisterService() {
	s.Router.GET("/api/distance-matrix", s.CalculateDistance)
}

// StartService will run the service on the desired port of the .env file
func (s Service) StartService(port string) {
	s.Router.Run(port)
}
