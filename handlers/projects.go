package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/svbrodersen/mySite/models"
	"github.com/svbrodersen/mySite/utils"
	"github.com/svbrodersen/mySite/views/templates"
)

type ProjectsDatabase struct {
	Db DataBaseAccessor
}

func (p *ProjectsDatabase) loggedInProjects(w http.ResponseWriter, r *http.Request) {
	db := p.Db
	if r.Method == "PATCH" {
		templates.ProjectList(true, true, db.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.ProjectList(false, true, db.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "POST" {
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error recieving file")
			templates.ProjectList(true, true, db.GetProjects()).Render(r.Context(), w)
			return
		}
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("error reading file")
			templates.ProjectList(true, true, db.GetProjects()).Render(r.Context(), w)
			return
		}
		html := utils.MdToHtml(fileBytes)
		title, content := utils.ExtractH1(html)
		project := models.Project{
			Title:       title,
			Description: r.FormValue("description"),
			Content:     content,
			Date:        time.Now().String(),
		}
		id, err := db.SaveProject(project)
		if err != nil {
			fmt.Println("error saving project")
		}
		fmt.Printf("Saved project as id: %d", id)
		templates.ProjectList(true, true, db.GetProjects()).Render(r.Context(), w)
	}
}

func (p ProjectsDatabase) defaultProjects(w http.ResponseWriter, r *http.Request) {
	db := p.Db
	if r.Method == "PATCH" {
		templates.ProjectList(true, false, db.GetProjects()).Render(r.Context(), w)
	} else {
		templates.ProjectList(false, false, db.GetProjects()).Render(r.Context(), w)
	}
}

func (p ProjectsDatabase) ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		p.defaultProjects(w, r)
		return
	}
	tknStr := cookie.Value
	err = utils.AuthVerifyToken(tknStr)
	// Verification failed
	if err != nil {
		p.defaultProjects(w, r)
		return
	}
	// Only admin logged in
	p.loggedInProjects(w, r)
}
