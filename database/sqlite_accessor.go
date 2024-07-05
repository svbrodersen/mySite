package database

import (
	"database/sql"
	"fmt"

	"github.com/svbrodersen/mySite/models"
	_ "modernc.org/sqlite"
)

type SqliteAccessor struct {
}

func (accessor SqliteAccessor) GetProjects() []models.Project {
	db, err := sql.Open("sqlite", "projects.db")
	if err != nil {
		panic("failed to open database")
	}
	rows, err := db.Query("SELECT * FROM projects;")
	if err != nil {
		fmt.Printf("Failed query: %s", err)
		return nil
	}
	defer db.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.Id,
			&project.Date,
			&project.Title,
			&project.Description,
			&project.Content,
		)
		projects = append(projects, project)
		if err != nil {
			fmt.Println("Failed scan")
			return projects
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Printf("Failed rows.err: %s", err)
		return projects
	}
	return projects
}

func (accessor SqliteAccessor) SaveProject(project models.Project) (int, error) {
	db, err := sql.Open("sqlite", "projects.db")
	if err != nil {
		panic("failed to open database")
	}

	res, err := db.Exec("INSERT INTO projects VALUES(NULL, ?, ?, ?, ?);",
		project.Date,
		project.Title,
		project.Description,
		project.Content,
	)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (accessor SqliteAccessor) GetProject(id int) (*models.Project, error) {
	db, err := sql.Open("sqlite", "projects.db")
	if err != nil {
		panic("failed to open database")
	}
	var project models.Project
	query_string := fmt.Sprintf("SELECT * FROM projects WHERE id=%d;", id)
	rows, err := db.Query(query_string)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows.Next()
	err = rows.Scan(
		&project.Id,
		&project.Date,
		&project.Title,
		&project.Description,
		&project.Content,
	)
	return &project, err
}
