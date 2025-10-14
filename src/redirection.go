package src

import (
	"html/template"
	"net/http"
	"strconv"

	"Power4-web/src"
)

var currentGame *src.Game

func main() {
	currentGame = src.NewGame()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/src", handleGame) // nouvelle route
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
	_ = tmpl.Execute(w, currentGame)
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		colStr := r.FormValue("column")
		col, err := strconv.Atoi(colStr)
		if err == nil {
			currentGame.PlayMove(col)
		}
	}
	http.Redirect(w, r, "/src", http.StatusSeeOther) // redirige maintenant vers la page du jeu
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	currentGame = src.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handleRestart(w http.ResponseWriter, r *http.Request) {
	currentGame = src.NewGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
