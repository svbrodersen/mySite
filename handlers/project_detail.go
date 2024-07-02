package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/svbrodersen/mySite/views/templates"
)

func (p *ProjectsDatabase) ProjectDetailHandler(w http.ResponseWriter, r *http.Request) {
	str_split := strings.Split(r.URL.String(), "/")
	str_id := str_split[len(str_split)-1]
	id, err := strconv.ParseInt(str_id, 10, 32)
	if err != nil {
		fmt.Printf("Error occured: %s", err)
		templates.Error404().Render(r.Context(), w)
	}
	db := p.Db
	project, err := db.GetProject(int(id))
	if err != nil {
		fmt.Printf("Error occured: %s", err)
		templates.Error404().Render(r.Context(), w)
	}
	if r.Method == "GET" {
		templates.ProjectDetail(false, project).Render(r.Context(), w)
	} else if r.Method == "PATCH" {
		templates.ProjectDetail(true, project).Render(r.Context(), w)
	}
}
