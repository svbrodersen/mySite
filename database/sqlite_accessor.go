package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/svbrodersen/mySite/models"
)

type SqliteAccessor struct {
	Db *sql.DB
}

func NewSqliteAccessor() (*SqliteAccessor, error) {
	db, err := sql.Open("sqlite3", "projects.db")
	return &SqliteAccessor{db}, err
}

func (accessor SqliteAccessor) GetProjects() []models.Project {
	db := accessor.Db
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

func (accessor SqliteAccessor) SaveProject(project models.Project) (int, error) {
	db := accessor.Db

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

func (accessor SqliteAccessor) GetProject(id int) (*models.Project, error) {
	db := accessor.Db
	var project models.Project
	query_string := fmt.Sprintf("SELECT * FROM projects WHERE id=%d;", id)
	rows, err := db.Query(query_string)
	if err != nil {
		return nil, err
	}
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
