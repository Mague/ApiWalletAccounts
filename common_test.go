package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Mague/ApiWalletAccounts/api"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	api.Account{}.Load(r)
	api.User{}.Load(r)
	api.Auth{}.Load(r)
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}
