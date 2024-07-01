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

func loggedInProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.ProjectList(true, true, utils.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.ProjectList(false, true, utils.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "POST" {
		fmt.Println("Uploading file")
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error recieving file")
			templates.ProjectList(true, true, utils.GetProjects()).Render(r.Context(), w)
			return
		}
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("error reading file")
			templates.ProjectList(true, true, utils.GetProjects()).Render(r.Context(), w)
			return
		}
		html := utils.MdToHtml(fileBytes)
		fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", fileBytes, html)
		title, content := utils.ExtractH1(html)
		project := models.Project{
			Title:       title,
			Description: r.FormValue("description"),
			Content:     content,
			Date:        time.Now().String(),
		}
		id, err := utils.SaveProject(project)
		if err != nil {
			fmt.Println("error saving project")
		}
		fmt.Printf("Saved project as id: %d", id)
		templates.ProjectList(true, true, utils.GetProjects()).Render(r.Context(), w)
	}
}

func defaultProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.ProjectList(true, false, utils.GetProjects()).Render(r.Context(), w)
	} else {
		templates.ProjectList(false, false, utils.GetProjects()).Render(r.Context(), w)
	}
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		defaultProjects(w, r)
		return
	}
	tknStr := cookie.Value
	err = utils.AuthVerifyToken(tknStr)
	// Verification failed
	if err != nil {
		defaultProjects(w, r)
		return
	}
	// Only admin logged in
	loggedInProjects(w, r)
}
