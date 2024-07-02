package handlers

import (
	"github.com/svbrodersen/mySite/models"
)

type DataBaseAccessor interface {
	GetProjects() []models.Project
	SaveProject(models.Project) (int, error)
	GetProject(int) (*models.Project, error)
}
