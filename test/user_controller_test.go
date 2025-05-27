package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yuskayuss/project-api/config"
	"github.com/yuskayuss/project-api/routes"
)

func TestCreateUser(t *testing.T) {
	config.ConnectDB()
	router := gin.Default()
	routes.UserRoute(router)

	body := `{"name":"Test","email":"test@mail.com","password":"test123"}`
	req, _ := http.NewRequest("POST", "/api/users/", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 but got %d", w.Code)
	}
}
