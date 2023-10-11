package main

import (
	"assignment3/controller"
	"assignment3/database"
	routers "assignment3/router"
	"net/http"
)

func main() {
	database.StartDB()
	db := database.GetDB()

	go controller.Update(db)
	http.ListenAndServe(":8080", routers.SetRouter())
}
