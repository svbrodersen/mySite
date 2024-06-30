package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/svbrodersen/mySite/models"
	"github.com/svbrodersen/mySite/utils"
	"github.com/svbrodersen/mySite/views/templates"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHtml(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func loggedInProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Projects(true, true, utils.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Projects(false, true, utils.GetProjects()).Render(r.Context(), w)
	} else if r.Method == "POST" {
		fmt.Println("Uploading file")
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error recieving file")
			templates.Projects(true, true, utils.GetProjects()).Render(r.Context(), w)
			return
		}
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("error reading file")
			templates.Projects(true, true, utils.GetProjects()).Render(r.Context(), w)
			return
		}
		html := fmt.Sprintf("%s", mdToHtml(fileBytes))
		fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", fileBytes, html)
		project := models.Project{
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			Content:     html,
			Date:        time.Now().String(),
		}
		id, err := utils.SaveProject(project)
		if err != nil {
			fmt.Println("Error saving project")
		}
		fmt.Println("Saved project as id: %s", id)
		templates.Projects(true, true, utils.GetProjects()).Render(r.Context(), w)
	}
}

func defaultProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Projects(true, false, utils.GetProjects()).Render(r.Context(), w)
	} else {
		templates.Projects(false, false, utils.GetProjects()).Render(r.Context(), w)
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
