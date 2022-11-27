package hangman

import (
	"net/http"
	"text/template"
)

type GameVars struct {
	IsStarted        bool
	Try              int
	Letter           string
	Mystery          string
	MystArr          []string
	Aleaindexonfiles string
	Stock            []string
	Stockstring      string
	Message          string
}

var game GameVars

func Hangman(w http.ResponseWriter, r *http.Request) {
	if (game.Aleaindexonfiles != "" && game.Try == 0) || game.IsStarted == false {
		game.IsStarted = true
		game.Try = 10
		game.Stock = []string{}
		game.Stockstring = ""
		game.Aleaindexonfiles = Pickword()
		game.Mystery = MysteryWord(game.Aleaindexonfiles)
	}

	RenderTemplate(w, "index.html")

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, game)

}
func Letter(w http.ResponseWriter, r *http.Request) {
	game.Message = ""
	game.Letter = r.FormValue("Letter")
	if ContainsInStock(game.Stock, game.Letter) {
		game.Message = "You already tried this letter"
	}
	if game.Message != "You already tried this letter" {

		if len(game.Letter) == 2 && Accents(&game.Letter) {
		}
		game = Stock(game.Letter, game)

		if len(game.Letter) == 1 {
			if Checkletter(game.Letter, game.Aleaindexonfiles) {
				game = Addletter(game.Letter, game.Aleaindexonfiles, game.Mystery, game)
				if Checkwin(game.Letter, game.Mystery, game.Aleaindexonfiles) {
					http.Redirect(w, r, "/win", http.StatusSeeOther)
				}
			} else {
				game.Try -= 1
			}
		} else if len(game.Letter) > 1 {
			if game.Letter == game.Aleaindexonfiles {
				http.Redirect(w, r, "/win", 302)
				return
			} else {
				game.Try -= 2

			}
		}
	}
	if Checklose(game.Try) {
		http.Redirect(w, r, "/lose", 302)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func Win(w http.ResponseWriter, r *http.Request) {

	if Checkwin(game.Letter, game.Mystery, game.Aleaindexonfiles) {
		RenderTemplate(w, "win.html")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func Lose(w http.ResponseWriter, r *http.Request) {
	if Checklose(game.Try) {
		RenderTemplate(w, "lose.html")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func Retry(w http.ResponseWriter, r *http.Request) {
	game.IsStarted = false
	http.Redirect(w, r, "/", 302)
}
