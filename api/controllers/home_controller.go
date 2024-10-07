package controllers

import (
	"GoLang-JWT-Postgres-Mysql-Restful-API/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
