package service

import "testing"

func TestNewService(t *testing.T) {
	expected := "API KEY GANZ"
	s := NewService(expected)

	if s.Key != expected {
		t.Fatal("Initializing Service Error")
	}
}
