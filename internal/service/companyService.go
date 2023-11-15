package service

import (
	"context"

	"job-portal-api/internal/models"
)

func (s *Service) AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error) {
	companyData, err := s.UserRepo.CreateCompany(ctx, companyData)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}

func (s *Service) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {
	companyDetails, err := s.UserRepo.ViewCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return companyDetails, nil
}

func (s *Service) ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error) {
	companyData, err := s.UserRepo.ViewCompanyById(ctx, cid)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
