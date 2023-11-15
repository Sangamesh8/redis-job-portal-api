package repository

import (
	"context"
	"errors"

	"job-portal-api/internal/models"

	"gorm.io/gorm"
)

// Repo struct represents the repository implementation with a reference to the database.
type Repo struct {
	DB *gorm.DB
}

// UserRepo is an interface defining methods for interacting with user-related, company-related, and job-related data.
//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository
type UserRepo interface {
	CreateUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)

	CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)

	CreateJob(ctx context.Context, jobData models.Jobs) (models.ResponseForJobs, error)
	FindJobByCompanyID(ctx context.Context, CompanyID uint64) ([]models.Jobs, error)
	FindAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobDetailsByJobId(ctx context.Context, jid uint64) (models.Jobs, error)
	JobByID(ctx context.Context, jid uint) (models.Jobs, error)
}
// NewRepository creates a new repository instance with the provided database connection.
func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
