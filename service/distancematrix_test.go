package service

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestService_CalculateDistance(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil && os.Getenv("KEY") == "" {
		log.Fatal("Error loading .env file :", err.Error())
	}
	app := NewService(os.Getenv("KEY"))

	app.RegisterService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/distance-matrix?origins=41.6655101,-73.8918896999998|40.6655101,-73.8918896999998&destinations=40.6905615%2C-73.9976592|42.6905615%2C-73.9976592", nil)
	app.Router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatal("Endpoint not returning 200")
	}
}
