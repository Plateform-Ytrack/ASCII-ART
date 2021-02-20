package main// Permet d'informer votre compilateur Go que le paquet doit être compilé en tant que programme exécutable au lieu d'un package utilitaire.

import (//Pour importer les packages
	"fmt"// Package fmt implémente les E/S formatées avec des fonctions analogues à celles de C printf et scanf.
	"io/ioutil"// Package io fournit des interfaces de base aux primitives d’E/S. Package ioutil implémente certaines fonctions utilitaires d’E/S.
	"os"// Package os fournit une interface indépendante de la plateforme pour les fonctionnalités du système d’exploitation.
)

func main() {// Les programmes go se lancent en exécutant une fonction appelé main
	arg := len(os.Args) // 
	if arg > 1 {//
 		texte(os.Args[1])//
	} 
} 

func texte(mot string) { //

	lire, fichier := ioutil.ReadFile("standard.txt") // Dans ce bloc, on demande de lire le fichier qui se situe dans les guillemets
	if fichier != nil { // On instaure une condition, dans laquelle si le fichier ne contient rien alors 
		fmt.Print(fichier) // Il nous renverra une erreur
	} 

		read := string(lire)
		standard := []rune(read)
		word := []rune(mot)
		var reult [10]string
		var ligne int = 0

		var ascii int = 32

		for _, c := range word {
			for i := range standard {
				if ascii == int((c + 1)) {
					break
				}
			}
		}
			if i < len(standard)-2 {//
				if standard[i] == 10 && standard[i+1] == 10 { //
					ascii++//
					ligne = 0 //On regarde à quelle lettre corresponds en incrémentant l’ASCII
				}
			}
			if int(c) == ascii {//
				if standard[i] == 10 {//
					ligne ++//
				} else if ligne < 10 {//
					result[ligne] = result[ligne] +//Il compare la valeur ASCII a int(c ) ou c est le caractère, c’est donc ici qu’on attribut les valeurs à result
				}
			}
		}
		ascii = 32//
	}
	for i := range result {//
		fmt.Println (result[i])//La boucle for imprime le contenu de result qui contient le shéma du caractère
	}
}
	