package service

import (
	"calc_example/internal/model"
	"calc_example/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Issue Service
func (s *Service) CreateIssue(req *model.CreateIssueRequest) (*model.IssueResponse, error) {
	issue := &model.Issue{
		FullName:              req.FullName,
		ContactInfo:           req.ContactInfo,
		PreferredContactMethod: req.PreferredContactMethod,
		HasChinaExperience:    req.HasChinaExperience,
		HasSupplierContacts:   req.HasSupplierContacts,
		ProductDescription:    req.ProductDescription,
		ExistingProductLinks:  req.ExistingProductLinks,
		Volume:                req.Volume,
		Weight:                req.Weight,
		Density:               req.Density,
		PreviousInvoiceFile:   req.PreviousInvoiceFile,
		ExpectedDeliveryDate:  req.ExpectedDeliveryDate,
		Status:                "open",
	}

	if err := s.repo.CreateIssue(issue); err != nil {
		return nil, err
	}

	return &model.IssueResponse{
		ID:                    issue.ID,
		FullName:              issue.FullName,
		ContactInfo:           issue.ContactInfo,
		PreferredContactMethod: issue.PreferredContactMethod,
		HasChinaExperience:    issue.HasChinaExperience,
		HasSupplierContacts:   issue.HasSupplierContacts,
		ProductDescription:    issue.ProductDescription,
		ExistingProductLinks:  issue.ExistingProductLinks,
		Volume:                issue.Volume,
		Weight:                issue.Weight,
		Density:               issue.Density,
		PreviousInvoiceFile:   issue.PreviousInvoiceFile,
		ExpectedDeliveryDate:  issue.ExpectedDeliveryDate,
		Status:                issue.Status,
		CreatedAt:             issue.CreatedAt,
		UpdatedAt:             issue.UpdatedAt,
	}, nil
}

func (s *Service) GetIssueByID(id uint) (*model.IssueResponse, error) {
	issue, err := s.repo.GetIssueByID(id)
	if err != nil {
		return nil, err
	}

	return &model.IssueResponse{
		ID:                    issue.ID,
		FullName:              issue.FullName,
		ContactInfo:           issue.ContactInfo,
		PreferredContactMethod: issue.PreferredContactMethod,
		HasChinaExperience:    issue.HasChinaExperience,
		HasSupplierContacts:   issue.HasSupplierContacts,
		ProductDescription:    issue.ProductDescription,
		ExistingProductLinks:  issue.ExistingProductLinks,
		Volume:                issue.Volume,
		Weight:                issue.Weight,
		Density:               issue.Density,
		PreviousInvoiceFile:   issue.PreviousInvoiceFile,
		ExpectedDeliveryDate:  issue.ExpectedDeliveryDate,
		Status:                issue.Status,
		CreatedAt:             issue.CreatedAt,
		UpdatedAt:             issue.UpdatedAt,
	}, nil
}

func (s *Service) GetAllIssues() ([]model.IssueResponse, error) {
	issues, err := s.repo.GetAllIssues()
	if err != nil {
		return nil, err
	}

	var responses []model.IssueResponse
	for _, issue := range issues {
		responses = append(responses, model.IssueResponse{
			ID:                    issue.ID,
			FullName:              issue.FullName,
			ContactInfo:           issue.ContactInfo,
			PreferredContactMethod: issue.PreferredContactMethod,
			HasChinaExperience:    issue.HasChinaExperience,
			HasSupplierContacts:   issue.HasSupplierContacts,
			ProductDescription:    issue.ProductDescription,
			ExistingProductLinks:  issue.ExistingProductLinks,
			Volume:                issue.Volume,
			Weight:                issue.Weight,
			Density:               issue.Density,
			PreviousInvoiceFile:   issue.PreviousInvoiceFile,
			ExpectedDeliveryDate:  issue.ExpectedDeliveryDate,
			Status:                issue.Status,
			CreatedAt:             issue.CreatedAt,
			UpdatedAt:             issue.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *Service) UpdateIssue(id uint, req *model.UpdateIssueRequest) (*model.IssueResponse, error) {
	issue, err := s.repo.GetIssueByID(id)
	if err != nil {
		return nil, err
	}

	issue.Status = req.Status

	if err := s.repo.UpdateIssue(issue); err != nil {
		return nil, err
	}

	return &model.IssueResponse{
		ID:                    issue.ID,
		FullName:              issue.FullName,
		ContactInfo:           issue.ContactInfo,
		PreferredContactMethod: issue.PreferredContactMethod,
		HasChinaExperience:    issue.HasChinaExperience,
		HasSupplierContacts:   issue.HasSupplierContacts,
		ProductDescription:    issue.ProductDescription,
		ExistingProductLinks:  issue.ExistingProductLinks,
		Volume:                issue.Volume,
		Weight:                issue.Weight,
		Density:               issue.Density,
		PreviousInvoiceFile:   issue.PreviousInvoiceFile,
		ExpectedDeliveryDate:  issue.ExpectedDeliveryDate,
		Status:                issue.Status,
		CreatedAt:             issue.CreatedAt,
		UpdatedAt:             issue.UpdatedAt,
	}, nil
}

func (s *Service) DeleteIssue(id uint) error {
	return s.repo.DeleteIssue(id)
} 