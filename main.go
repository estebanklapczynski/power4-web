package main

import (
	"log"
	"net/http"
	"power4-web/src/web"
)

func main() {
	// Configure toutes les routes
	web.SetupRoutes()

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
