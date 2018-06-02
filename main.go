package main

import (
	"./models"
	"./routes"
	"./utils"
	"net/http"
)

func main() {
	models.Init()
	utils.LoadTemplate("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
