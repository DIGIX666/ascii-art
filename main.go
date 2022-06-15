package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getTable() []string {
	file, err := os.Open("standard.txt")
	content, _ := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	table := strings.Split(string(content), "\n")

	return table
}
func asciiGenerator(text string) {

	line := strings.Split(text, "\\n")
	content := getTable()
	for i := 0; i < len(line); i++ {
		if len(line[i]) > 0 {
			chars := []rune(line[i])
			for n := 0; n < 8; n++ { // vérifie qu'on ne dépasse pas la hauteur de 8 lignes pour les caracts et lettres 
				for v := 0; v < len(chars); v++ {
					group := int(chars[v]) - 32
					adress := group * 9  // permet de récupper le bon cararct et de les grouper,
					charLine := content[adress+1+n]
					fmt.Print(charLine)
				}
				fmt.Print(string(rune('\n')))
			}
		} else {
			fmt.Print(string(rune('\n')))
		}
	}
}

func main() {
	args := os.Args
	if len(args) == 2 { // car dans la ligne de commande ils nous faut juste le nom du fichier + l'argument 
		asciiGenerator(os.Args[1])
	} else {
		fmt.Println("Erreur  nombre d'arguments")
		fmt.Println("Usage exemple : go run . \"text\"")
	}
}
