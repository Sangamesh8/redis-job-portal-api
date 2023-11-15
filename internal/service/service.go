package service

import (
	"context"
	"errors"

	"job-portal-api/internal/auth"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     auth.Authentication
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
type JobPortalService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserSignIn(ctx context.Context, userData models.NewUser) (string, error)

	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error)
	ViewJobByCompanyID(ctx context.Context, cid uint64) ([]models.Jobs, error)

	AddJobDetails(ctx context.Context, jobData models.CreateJobs, CompanyID uint64) (models.ResponseForJobs, error)
	ViewAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error)

	ProcessJobApplication(ctx context.Context, jobData []models.JobApplicantResponse) ([]models.JobApplicantResponse, error)
}

func NewService(userRepo repository.UserRepo, a auth.Authentication) (JobPortalService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
	}, nil
}
