package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"power4-web/game"
)

var new_game *game.Game

func main() {
	new_game = game.NewGame()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, new_game)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		colStr := r.FormValue("column")
		col, err := strconv.Atoi(colStr)
		if err == nil {
			new_game.PlayMove(col) // Correction ici
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	new_game = game.NewGame() // Correction ici
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
