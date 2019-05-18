package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// handler is http struct for passing services to the router.
type handler struct {
	Logger *log.Logger
}

// goHome is used for handling requests to "/".
func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "page-template-split"}
	respond(h.Logger, w, r, "./tmpl/index.tmpl", data)
}

// contact is used for handling requests to "/contact".
func (h *handler) contact(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "page-template-default"}
	respond(h.Logger, w, r, "./tmpl/contact.tmpl", data)
}

// caseStudies is used for handling requests to "/case-studies".
func (h *handler) caseStudies(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "page-template-carousel"}
	respond(h.Logger, w, r, "./tmpl/case_studies.tmpl", data)
}

// caseStudyShow is used for handling requests to "/case-study/:id".
func (h *handler) caseStudyShow(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"bodyClass": "single single-portfolio"}
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	respond(h.Logger, w, r, "./tmpl/case_study_"+id+".tmpl", data)
}

// respond is used to parse a base template.
func respond(logger *log.Logger, w http.ResponseWriter, r *http.Request, layout string, data interface{}) {
	// Parse static files.
	tmpl := template.Must(template.New("base.tmpl").Funcs(templateFuncs).ParseFiles(
		"./tmpl/base.tmpl",
		"./tmpl/partials/_left_nav.tmpl",
		"./tmpl/partials/_right_sidebar.tmpl",
		"./tmpl/partials/_my_info.tmpl",
		// "./tmpl/partials/_menu_hero.tmpl",
		// "./tmpl/partials/_menu.tmpl",
		// "./tmpl/partials/_drinks_hero.tmpl",
		// "./tmpl/partials/_drinks.tmpl",
		// "./tmpl/partials/_wall.tmpl",
		// "./tmpl/partials/_owner_hero.tmpl",
		// "./tmpl/partials/_kyle.tmpl",
		// "./tmpl/partials/_ramon.tmpl",
		// "./tmpl/partials/_catering.tmpl",
		// "./tmpl/partials/_map.tmpl",
		// "./tmpl/partials/_footer.tmpl",
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
