package controllers

import (
	"net/http"

	"github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}