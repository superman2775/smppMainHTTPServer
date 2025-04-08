package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"smppmainhttpserver/utils"
	"strings"
	"time"
)

func SetModeHandler(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	redirect := r.URL.Query().Get("redirect")
	if mode != "light" && mode != "dark" {
		mode = "light"
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "mode",
		Value:   mode,
		Expires: time.Now().Add(365 * 24 * time.Hour),
	})

	if strings.Contains(redirect, ".") {
		return
	}
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}
func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles("content/template/template.html", tmpl)
	if err != nil {
		http.Error(w, "404 Not Found in renderTemplate", http.StatusNotFound)
		log.Printf("Failed to parse templates: %v, error: %v", tmpl, err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}
func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	pageTitle := "Title"
	pageName := page
	if page == "" {
		page = "index"
		pageName = "Home"
		pageTitle = "Smartschool++"
	} else {
		pageTitle = utils.ToUpperCase(page)
	}

	w.Header().Set("Accept-CH", "Sec-CH-Prefers-Color-Scheme")
	theme := r.Header.Get("Sec-CH-Prefers-Color-Scheme")
	if theme != "" {
		theme = "dark"
	}
	cookie, err := r.Cookie("mode")
	if err == nil {
		theme = cookie.Value
	}

	cssMainFile := page + ".css"
	cssMainFilePath := filepath.Join("content", "css", cssMainFile)
	if _, err := os.Stat(cssMainFilePath); os.IsNotExist(err) {
		cssMainFile = "index.css"
	}

	cssThemeFile := "light.css"
	if theme == "dark" {
		cssThemeFile = "dark.css"
	}

	tmplPath := filepath.Join("content", "html", page+".html")

	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		page = "404"
		pageTitle = "404 Page Not Found"
	}

	data := struct {
		Title     string
		ThemeCSS  string
		MainCSS   string
		Redirect  string
		PageTitle string
		Mode      string
	}{
		Title:     pageName,
		ThemeCSS:  cssThemeFile,
		MainCSS:   cssMainFile,
		Redirect:  r.URL.Path,
		PageTitle: pageTitle,
		Mode:      theme,
	}

	RenderTemplate(w, tmplPath, data)
}
