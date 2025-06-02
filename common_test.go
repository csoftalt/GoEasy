package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// article represents a simple article structure for testing purposes
var tmpArticleList []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// getRouter returns a new Gin engine with or without template loading
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}

	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Errorf("Expected response to match condition, got %d", w.Code)
	}
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
  tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
  articleList = tmpArticleList
}