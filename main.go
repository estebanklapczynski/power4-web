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
<<<<<<< HEAD
	currentGame = game.NewGame()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/game", handleGame)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)
	http.HandleFunc("/restart", handleRestart)
=======
	http.HandleFunc("/", serveStart)
	http.HandleFunc("/game", serveGame)
	http.HandleFunc("/play", playMove)
>>>>>>> 1fe4eba99b3f469bb529b2ad647312a742b204d9
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

<<<<<<< HEAD
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	tmpl.Execute(w, currentGame)
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	_ = tmpl.Execute(w, currentGame)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
=======
func serveStart(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/start.html"))
	_ = t.Execute(w, nil)
}

func serveGame(w http.ResponseWriter, _ *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	_ = t.Execute(w, g)
}

func playMove(w http.ResponseWriter, r *http.Request) {
>>>>>>> 1fe4eba99b3f469bb529b2ad647312a742b204d9
	if r.Method == http.MethodPost {
		if col, err := strconv.Atoi(r.FormValue("column")); err == nil {
			g.PlayMove(col)
		}
	}
<<<<<<< HEAD
	http.Redirect(w, r, "/game", http.StatusSeeOther) // redirige maintenant vers la page du jeu
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handleRestart(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
=======
	http.Redirect(w, r, "/game", http.StatusSeeOther)
>>>>>>> 1fe4eba99b3f469bb529b2ad647312a742b204d9
}
