package art

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web/ascii"
)

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		MethodNotAllowed(w, r)
		return
	}

	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/home_page.html", "templates/header.html")
		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}

		err = t.ExecuteTemplate(w, "home", nil)
		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}

	}
}

// Asciiart ...
func Asciiart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		MethodNotAllowed(w, r)
		return
	}

	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/home_page.html", "templates/header.html")
		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}
		err = t.ExecuteTemplate(w, "home", nil)
		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}
	}

	if r.Method == http.MethodPost {
		_, err := template.ParseFiles("templates/home_page.html", "templates/header.html")
		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}

		Word := r.FormValue("text")
		if Word == "" {
			BadRequest(w, r)
			return
		}

		if err = ascii.CheckValid(Word); err != nil {
			BadRequest(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}

		Banner := r.FormValue("banner")
		if err = ascii.CheckFile(Banner); err != nil {
			BadRequest(w, r)
			return
		}

		if err != nil {
			InternalServerError(w, r)
			fmt.Fprintf(w, err.Error())
			log.Println("\n" + err.Error())
			return
		}
	}
}
