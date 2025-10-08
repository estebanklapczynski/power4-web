package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"power4-web/game"
)

var g = game.NewGame()

func main() {
	http.HandleFunc("/", serveStart)
	http.HandleFunc("/game", serveGame)
	http.HandleFunc("/play", playMove)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveStart(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/start.html"))
	_ = t.Execute(w, nil)
}

func serveGame(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	_ = t.Execute(w, g)
}

func playMove(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if col, err := strconv.Atoi(r.FormValue("column")); err == nil {
			g.PlayMove(col)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
