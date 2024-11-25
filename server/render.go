package server

import (
	"html/template"
	"log"
	"net/http"

	"web/generator"
)

// Render the main HTML template with the provided data.
func TemplateRender(w http.ResponseWriter, banners []string, inputString, asciiArt, color, backgroundcolor, banner string) {
	if color == "" {
		color = "#00ff99"
	}
	if backgroundcolor == "" {
		backgroundcolor = "#1f1f1f"
	}
	if banner == "" {
		banner = "standard"
	}
	// Prepare data for the HTML template.
	data := struct {
		AsciiArt        string
		Banners         []string
		Input           string
		Color           string
		BackgroundColor string
		Banner          string
	}{
		AsciiArt:        asciiArt,
		Banners:         banners,
		Input:           inputString,
		Color:           color,
		BackgroundColor: backgroundcolor,
		Banner:          banner,
	}

	// Parse and execute the HTML template.
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		log.Println("Template parsing error: ", err)
		return
	}

	tmpl.Execute(w, data)
}

// Requests handling for the main page and processes form submissions.
func PageRender(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusNotFound)
		return
	}

	banners, err := generator.BannerList()
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		log.Println("Banner list error: ", err)
		return
	}

	var inputString, asciiArt, color, backgroundcolor, banner string
	// Handle POST request: process form inputs and generate ASCII art.
	if r.Method == http.MethodPost {
		inputString = r.FormValue("input")
		banner = r.FormValue("banner")
		color = r.FormValue("color")
		backgroundcolor = r.FormValue("background-color")

		if inputString == "" {
			ErrorPage(w, http.StatusBadRequest)
			return
		}
		asciiArt, err = generator.PrintAsciiArt(inputString, banner)
		if err != nil {
			ErrorPage(w, http.StatusBadRequest)
			log.Print("ASCII art generation error: ", err)
			return
		}
	} else if r.Method == http.MethodGet {
		// Handle GET request: render the template with default values.
		TemplateRender(w, banners, "", "", "", "", "")
		return
	} else {
		ErrorPage(w, http.StatusMethodNotAllowed)
	}
	// Render the template with the processed data.
	TemplateRender(w, banners, inputString, asciiArt, color, backgroundcolor, banner)
}
