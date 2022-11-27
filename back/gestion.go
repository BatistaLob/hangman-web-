package hangman

import (
	"strings"
)

func ContainsInStock(word []string, letter string) bool {
	for _, c := range word {
		if c == letter {
			return true
		}
	}
	return false
}

func Checkletter(l string, aleaindexonfiles string) bool {
	if strings.Contains(aleaindexonfiles, l) {
		return true
	}
	return false
}

func Checkstring(l string, word string) bool {
	if l == word {
		return true
	}
	return false
}

func Accents(letter *string) bool {
	accents := map[string][]string{
		"e": {"é", "è", "ê", "ë"},
		"a": {"à", "â", "ä"},
		"i": {"î", "ï"},
		"o": {"ô", "ö"},
		"u": {"ù", "û", "ü"},
		"c": {"ç"},
	}
	for k, v := range accents {
		for _, c := range v {
			if *letter == c {
				*letter = k
				return true
			}
		}
	}
	return false
}
func StringInArray(s string) []string {
	tabstring := []string{}
	for _, c := range s {
		tabstring = append(tabstring, string(c))
	}
	return tabstring
}
func Stock(letter string, game GameVars) GameVars {
	if !ContainsInStock(game.Stock, letter) {
		game.Stock = append(game.Stock, letter)
	}
	game.Stockstring = strings.Join(game.Stock, " ")
	return game
}
