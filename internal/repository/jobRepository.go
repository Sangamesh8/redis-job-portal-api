package repository

import (
	"context"
	"errors"
	"fmt"

	"job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) ViewJobDetailsByJobId(ctx context.Context, jid uint64) (models.Jobs, error) {
	var jobData models.Jobs
	result := r.DB.Preload("JobLocation").
		Preload("Technology").
		Preload("WorkMode").
		Preload("Qualification").
		Preload("Shift").
		Preload("JobType").
		Where("id = ?", jid).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, errors.New("could not create the jobs")
	}
	fmt.Println("Body data -----", jobData)
	return jobData, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Jobs) (models.ResponseForJobs, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.ResponseForJobs{}, errors.New("could not create the jobs")
	}
	return models.ResponseForJobs{
		ID: jobData.ID,
	}, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Jobs, error) {
	var jobDatas []models.Jobs
	result := r.DB.Find(&jobDatas)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the jobs")
	}
	return jobDatas, nil

}

func (r *Repo) FindJobByCompanyID(ctx context.Context, CompanyID uint64) ([]models.Jobs, error) {
	var jobData []models.Jobs
	result := r.DB.Where("company_id = ?", CompanyID).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the company")
	}
	return jobData, nil
}

func (r *Repo) JobByID(ctx context.Context, jid uint) (models.Jobs, error) {
	var jobData models.Jobs
	result := r.DB.Preload("JobLocation").
		Preload("Technology").
		Preload("WorkMode").
		Preload("Qualification").
		Preload("Shift").
		Preload("JobType").
		Where("id = ?", jid).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, errors.New("couldn't process the job data")
	}
	return jobData, nil
}
