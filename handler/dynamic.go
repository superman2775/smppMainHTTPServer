package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"smppmainhttpserver/i18n"
	"smppmainhttpserver/utils"
	"strings"
	"time"

	goI18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func SetModeHandler(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	redirect := r.URL.Query().Get("redirect")
	if mode != "light" && mode != "dark" {
		mode = "dark"
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

func SetLanguageHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	redirect := r.URL.Query().Get("redirect")
	if lang != "en" && lang != "nl" {
		lang = "en"
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "lang",
		Value:   lang,
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
		log.Printf("Failed to parse templates: %v, error: %v", tmpl, err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}
func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	page := strings.ToLower(r.URL.Path[1:])
	pageTitle := "Title"
	if page == "" {
		page = "index"
		pageTitle = "Smartschool++"
	} else {
		pageTitle = utils.RemoveDash(page)
		pageTitle = utils.ToUpperCase(pageTitle)
	}

	// Get language preference
	lang := "en"
	cookie, err := r.Cookie("lang")
	if err == nil {
		lang = cookie.Value
	} else {
		// Fallback to browser language
		acceptLang := r.Header.Get("Accept-Language")
		if strings.Contains(acceptLang, "nl") {
			lang = "nl"
		}
	}
	localizer := i18n.Localizer(lang)

	w.Header().Set("Accept-CH", "Sec-CH-Prefers-Color-Scheme")
	userAgent := r.Header.Get("User-Agent")
	isMobile := strings.Contains(strings.ToLower(userAgent), "mobile") ||
		strings.Contains(strings.ToLower(userAgent), "android") ||
		strings.Contains(strings.ToLower(userAgent), "iphone") ||
		strings.Contains(strings.ToLower(userAgent), "ipad") ||
		strings.Contains(strings.ToLower(userAgent), "windows phone")

	theme := r.Header.Get("Sec-CH-Prefers-Color-Scheme")
	if theme == "" {
		theme = "dark"
	}
	modeCookie, modeErr := r.Cookie("mode")
	if modeErr == nil {
		theme = modeCookie.Value
	}

	jsMainFile := page + ".js"
	jsMainFilePath := filepath.Join("content", "js", jsMainFile)
	if _, err := os.Stat(jsMainFilePath); os.IsNotExist(err) {
		jsMainFile = "null"
	}

	cssMainFile := page + ".css"
	cssMainFilePath := filepath.Join("content", "css", cssMainFile)
	if _, err := os.Stat(cssMainFilePath); os.IsNotExist(err) {
		cssMainFile = "null"
	}

	cssThemeFile := "light.css"
	if theme == "dark" {
		cssThemeFile = "dark.css"
	}

	tmplPath := filepath.Join("content", "html", page+".html")

	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		page = "404"
		pageTitle = "404 Page Not Found"
		tmplPath = filepath.Join("content", "html", "404.html")
	}

	var markdownHTML template.HTML

	// Try to load localized markdown first
	markdownPath := filepath.Join("content", "md", lang, page+".md")
	if _, err := os.Stat(markdownPath); err == nil {
		mdTextBytes, readErr := os.ReadFile(markdownPath)
		if readErr != nil {
			http.Error(w, "Failed to read markdown", http.StatusInternalServerError)
			log.Println("Error reading "+page+".md for language "+lang+":", readErr)
			return
		}
		// Parse the markdown to HTML
		markdownHTML = template.HTML(utils.ParseMd(string(mdTextBytes), page)) // Mark it as raw HTML
	} else {
		// Fallback to English markdown
		markdownPath = filepath.Join("content", "md", page+".md")
		if _, err := os.Stat(markdownPath); err == nil {
			mdTextBytes, readErr := os.ReadFile(markdownPath)
			if readErr != nil {
				http.Error(w, "Failed to read markdown", http.StatusInternalServerError)
				log.Println("Error reading "+page+".md:", readErr)
				return
			}
			// Parse the markdown to HTML
			markdownHTML = template.HTML(utils.ParseMd(string(mdTextBytes), page)) // Mark it as raw HTML
		}
	}
	// Prepare JavaScript translations
	jsTranslations := map[string]string{
		"first_break":        i18n.TranslateString(localizer, "js_first_break", nil),
		"lunch":             i18n.TranslateString(localizer, "js_lunch", nil),
		"schools_out":       i18n.TranslateString(localizer, "js_schools_out", nil),
		"night_checkin":     i18n.TranslateString(localizer, "js_night_checkin", nil),
		"break_mobile":      i18n.TranslateString(localizer, "js_break_mobile", nil),
		"lunch_mobile":      i18n.TranslateString(localizer, "js_lunch_mobile", nil),
		"done_mobile":       i18n.TranslateString(localizer, "js_done_mobile", nil),
		"night_mobile":      i18n.TranslateString(localizer, "js_night_mobile", nil),
		"sweater_weather":   i18n.TranslateString(localizer, "js_sweater_weather", nil),
		"lessons_afternoon": i18n.TranslateString(localizer, "js_lessons_afternoon", nil),
		"your_bus_arrive":   i18n.TranslateString(localizer, "js_your_bus_arrive", nil),
		"homework_check":    i18n.TranslateString(localizer, "js_homework_check", nil),
		"weather_title":     i18n.TranslateString(localizer, "js_weather_title", nil),
		"planner_title":     i18n.TranslateString(localizer, "js_planner_title", nil),
		"delijn_title":      i18n.TranslateString(localizer, "js_delijn_title", nil),
		"volume_turn_on":    i18n.TranslateString(localizer, "js_volume_turn_on", nil),
		"volume_continue":   i18n.TranslateString(localizer, "js_volume_continue", nil),
	}

	data := struct {
		Page             string
		ThemeCSS         string
		MainCSS          string
		MainJS           string
		Redirect         string
		PageTitle        string
		Mode             string
		IsMobile         bool
		MarkdownHTML     template.HTML
		Lang             string
		JSTranslations   map[string]string
	}{
		Page:           page,
		ThemeCSS:       cssThemeFile,
		MainCSS:        cssMainFile,
		MainJS:         jsMainFile,
		Redirect:       r.URL.Path,
		PageTitle:      pageTitle,
		Mode:           theme,
		IsMobile:       isMobile,
		MarkdownHTML:   markdownHTML,
		Lang:           lang,
		JSTranslations: jsTranslations,
	}

	// Render template with proper localization
	renderTemplateWithLocalization(w, tmplPath, data, localizer)
}

func renderTemplateWithLocalization(w http.ResponseWriter, tmplPath string, data interface{}, localizer *goI18n.Localizer) {
	// Create FuncMap
	funcMap := template.FuncMap{
		"Localize": func(key string) string {
			text, err := i18n.Translate(localizer, key, nil)
			if err != nil {
				return key // Fallback to key if translation not found
			}
			return text
		},
	}

	// Create a new template with FuncMap
	tmpl := template.New("").Funcs(funcMap)
	
	// Read and parse main template
	mainContent, err := os.ReadFile("content/template/template.html")
	if err != nil {
		log.Printf("Failed to read main template: %v", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	
	// Parse main template
	tmpl, err = tmpl.Parse(string(mainContent))
	if err != nil {
		log.Printf("Failed to parse main template: %v", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	
	// Read and parse page template
	pageContent, err := os.ReadFile(tmplPath)
	if err != nil {
		log.Printf("Failed to read page template: %v", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	
	// Parse page template and add to collection
	_, err = tmpl.Parse(string(pageContent))
	if err != nil {
		log.Printf("Failed to parse page template: %v", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Execute the template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
