package hangman

import (
	"math/rand"
	"strings"
	"time"
)

func MysteryWord(aleaindexonfiles string) string {
	numberOfRand := (len(aleaindexonfiles) / 2) - 1
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(aleaindexonfiles))
	mysterytab := []string{}
	var mystery string
	for long := 0; long < len(aleaindexonfiles); long++ {
		mysterytab = append(mysterytab, "_")
	}
	for alea := 0; alea <= numberOfRand; {
		if mysterytab[random] == "_" {
			mysterytab[random] = string(aleaindexonfiles[random])
			random = rand.Intn(len(aleaindexonfiles))
			alea += 1
		} else {
			random = rand.Intn(len(aleaindexonfiles))
		}
	}
	for _, index := range mysterytab {
		mystery = mystery + index
	}
	return mystery
}
func Checkwin(letter string, mystery string, aleaindexonfiles string) bool {
	if len(letter) > 1 && letter == aleaindexonfiles {
		return true
	}
	if mystery == aleaindexonfiles {
		return true
	}
	return false
}

func Checklose(try int) bool {
	if try <= 0 {
		return true
	}
	return false
}
func Addletter(letter string, aleaindexonfiles string, mystery string, game GameVars) GameVars {
	game.MystArr = StringInArray(mystery)
	for ind, letMyst := range aleaindexonfiles {
		if string(letMyst) == letter {
			game.MystArr[ind] = letter
			game.Mystery = strings.Join(game.MystArr, "")
		}
	}
	return game
}
