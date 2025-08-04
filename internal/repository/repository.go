package repository

import (
	"calc_example/internal/model"
	"calc_example/pkg/database"
)

type Repository struct {
	db *database.Database
}

func New(db *database.Database) *Repository {
	return &Repository{db: db}
}

// Issue Repository
func (r *Repository) CreateIssue(issue *model.Issue) error {
	return r.db.Create(issue).Error
}

func (r *Repository) GetIssueByID(id uint) (*model.Issue, error) {
	var issue model.Issue
	err := r.db.First(&issue, id).Error
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

func (r *Repository) GetAllIssues() ([]model.Issue, error) {
	var issues []model.Issue
	err := r.db.Order("created_at DESC").Find(&issues).Error
	return issues, err
}

func (r *Repository) UpdateIssue(issue *model.Issue) error {
	return r.db.Save(issue).Error
}

func (r *Repository) DeleteIssue(id uint) error {
	return r.db.Delete(&model.Issue{}, id).Error
}
