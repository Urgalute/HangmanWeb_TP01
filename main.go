package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type DataPromoPage struct {
	Classe      string
	Filiere     string
	Niveau      string
	StudentQty  int
	StudentList []Student
}
type Student struct {
	Nom    string
	Prenom string
	Age    int  
	Sexe   string
}
type DataChangePage struct {
	Compteur int
	Texte    string
}

var Compteur int

func main() {

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		temp, tempErr := template.ParseFiles("./templates/index.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates")
			os.Exit(1)
		}
		datapage := DataPromoPage{
			Classe:      "B1 Cyber",
			Filiere:     "Cybersecurité",
			Niveau:      "Bachelor 1",
			StudentQty:  30,
			StudentList: []Student{{"John", "Doe", 20, "Homme"}, {"Jane", "Smith", 22, "Femme"}, {"Alice", "Williams", 21, "Femme"}},
		}
		temp.ExecuteTemplate(w, "index", datapage)
	})

	// Handle change page request and display the count of views and an appropriate message
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		temp, tempErr := template.ParseFiles("./templates/change.html")
		if tempErr != nil {
			fmt.Println("Error parsing templates")
			os.Exit(1)
		}
		Compteur++
		var Texte string
		if Compteur%2 == 0 {
			Texte = "Le nombre de vues est pair: "
		} else {
			Texte = "Le nombre de vues est impair: "
		}
		datapage := DataChangePage{
			Compteur: Compteur,
			Texte:    Texte,
		}
		temp.ExecuteTemplate(w, "change", datapage)
	})

	// Start server on port 8080
	fmt.Println("(http://localhost:8080/) - Server started on port:8000")

	http.ListenAndServe("localhost:8080", nil)

	Compteur = 0
}
