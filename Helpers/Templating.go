package Helpers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/foolin/goview"
)

func RenderTemplate(page uint, w http.ResponseWriter, r *http.Request) {
	gv := goview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"layouts/footer", "layouts/sidebar", "layouts/content"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"}})

	goview.Use(gv)

}
