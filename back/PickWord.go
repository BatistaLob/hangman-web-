package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ReadFileName(name string) string {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("ERROR: open " + name + ": no such file or directory\n")
		os.Exit(1)
	}
	res := ""
	arr := make([]byte, 1000)
	n, _ := file.Read(arr)
	for i := 0; i < n; i++ {
		res += string(arr[i])
	}
	return res

}
func ReadWordsOnFiles(wordsn string) []string {
	tab := strings.Split(wordsn, "\r\n")
	return tab
}
func Pickword() string {
	rand.Seed(time.Now().UnixNano())
	fileName := ReadFileName(os.Args[1])
	tabword := ReadWordsOnFiles(fileName)
	aleaindexonfiles := tabword[rand.Intn(len(tabword))]
	return aleaindexonfiles
}
