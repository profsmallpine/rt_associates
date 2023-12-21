package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// handler is http struct for passing services to the router.
type handler struct {
	Logger *log.Logger
}

// goHome is used for handling requests to "/".
func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "page-template-carousel"}
	respond(h.Logger, w, r, "./tmpl/index.tmpl", data)
}

// contact is used for handling requests to "/contact".
// func (h *handler) contact(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]interface{}{"bodyClass": "page-template-default"}
// 	respond(h.Logger, w, r, "./tmpl/contact.tmpl", data)
// }

// portfolioRose is used for handling requests to "/portfolio-rose".
func (h *handler) portfolioRose(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "single single-post"}
	respond(h.Logger, w, r, "./tmpl/portfolio_rose.tmpl", data)
}

// portfolioJim is used for handling requests to "/portfolio-jims".
func (h *handler) portfolioJim(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "single single-post"}
	respond(h.Logger, w, r, "./tmpl/portfolio_jim.tmpl", data)
}

// portfolioJim is used for handling requests to "/portfolio-jims".
func (h *handler) portfolioOptimum(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "single single-post"}
	respond(h.Logger, w, r, "./tmpl/portfolio_optimum.tmpl", data)
}

// respond is used to parse a base template.
func respond(logger *log.Logger, w http.ResponseWriter, r *http.Request, layout string, data interface{}) {
	// Parse static files.
	tmpl := template.Must(template.New("base.tmpl").Funcs(templateFuncs).ParseFiles(
		"./tmpl/base.tmpl",
		"./tmpl/partials/_left_nav.tmpl",
		"./tmpl/partials/_right_sidebar.tmpl",
		"./tmpl/partials/_my_info.tmpl",
		layout,
	))
	err := tmpl.Funcs(template.FuncMap{}).Execute(w, data)

	// Log template compilation failure.
	if err != nil {
		logger.Println("Template execution error: ", err.Error(), layout)
		return
	}
}

var templateFuncs = map[string]interface{}{
	"javascriptTag": javascriptTag,
	"stylesheetTag": stylesheetTag,
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
}
