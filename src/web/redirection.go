package web

import (
	"html/template"
	"net/http"
	"strconv"

	"power4-web/src/game"
)

var currentGame *game.Game

func SetupRoutes() {
	currentGame = game.NewGame()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/game", handleGame) // nouvelle route
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)
	http.HandleFunc("/restart", handleRestart)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Page1.html"))
	tmpl.Execute(w, currentGame)
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, currentGame)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		colStr := r.FormValue("column")
		col, err := strconv.Atoi(colStr)
		if err == nil {
			currentGame.PlayMove(col)
		}
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther) // redirige maintenant vers la page du jeu
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handleRestart(w http.ResponseWriter, r *http.Request) {
	currentGame = game.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
