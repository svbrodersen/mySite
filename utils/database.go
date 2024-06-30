package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/svbrodersen/mySite/models"
)

func GetProjects() []models.Project {
	db, err := sql.Open("sqlite3", "projects.db")
	if err != nil {
		fmt.Printf("Failed open: %s", err)
		return nil
	}

	rows, err := db.Query("SELECT * FROM projects;")
	if err != nil {
		fmt.Printf("Failed query: %s", err)
		return nil
	}

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

func SaveProject(project models.Project) (int, error) {
	db, err := sql.Open("sqlite3", "projects.db")
	if err != nil {
		return 0, err
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

	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
