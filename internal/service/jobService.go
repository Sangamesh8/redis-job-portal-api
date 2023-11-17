package service

import (
	"context"
	"encoding/json"
	"fmt"
	"job-portal-api/internal/models"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (s *Service) AddJobDetails(ctx context.Context, jobData models.CreateJobs, CompanyID uint64) (models.ResponseForJobs, error) {
	job := models.Jobs{
		CompanyID:       uint(CompanyID),
		Name:            jobData.Name,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,
		Budget:          jobData.Budget,
		JobDescription:  jobData.JobDescription,
	}

	for _, v := range jobData.JobLocation {
		jobLoc := models.JobLocation{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobLocation = append(job.JobLocation, jobLoc)
	}
	for _, v := range jobData.Technology {
		jobTec := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Technology = append(job.Technology, jobTec)
	}
	for _, v := range jobData.Qualification {
		jobQual := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Qualification = append(job.Qualification, jobQual)
	}
	for _, v := range jobData.Shift {
		jobShift := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Shift = append(job.Shift, jobShift)
	}
	for _, v := range jobData.JobType {
		jobJobType := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobType = append(job.JobType, jobJobType)
	}
	for _, v := range jobData.WorkMode {
		jobworkMode := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.WorkMode = append(job.WorkMode, jobworkMode)
	}
	createdJob, err := s.UserRepo.CreateJob(ctx, job)
	if err != nil {
		return models.ResponseForJobs{}, err
	}
	return createdJob, nil
}

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {

	jobData, err := s.UserRepo.ViewJobDetailsByJobId(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) ViewJobByCompanyID(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FindJobByCompanyID(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (s *Service) ProcessJobApplication(ctx context.Context, jobData []models.JobApplicantResponse) ([]models.JobApplicantResponse, error) {
	var ProccessedJobData []models.JobApplicantResponse

	ch := make(chan models.JobApplicantResponse) // make a channel
	wg := new(sync.WaitGroup)                    // Initialize waitgroup variable

	for _, v := range jobData {
		wg.Add(1)                                // increment the waitgroup variable
		go func(v models.JobApplicantResponse) { // goroutine
			defer wg.Done() // decrement the waitgroup variable
			var jobDa models.Jobs
			val, err := s.rdb.Get(ctx, v.JobID)
			fmt.Println("Redis 1", err)
			if err == redis.Nil {

				dbData, err := s.UserRepo.JobByID(ctx, v.JobID)
				if err != nil {
					return
				}
				err = s.rdb.Set(ctx, v.JobID, dbData)
				if err != nil {
					return
				}
				jobDa = dbData
			} else {
				err = json.Unmarshal([]byte(val), &jobDa)
				if err == redis.Nil {
					return
				}
				if err != nil {
					return
				}
			}
			check, _ := applicationFilter(v, jobDa)
			if check {
				ch <- v
			}
		}(v)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for data := range ch {
		ProccessedJobData = append(ProccessedJobData, data)
	}
	return ProccessedJobData, nil
}

func applicationFilter(validateApplication models.JobApplicantResponse, jobDetails models.Jobs) (bool, models.JobApplicantResponse) {
	// convert string to integer and store in applicantBudget variable
	applicantBudget, err := strconv.Atoi(validateApplication.Jobs.Budget)
	if err != nil {
		panic("error while conversion budget data from applicants")
	}
	// convert string to integer and store in compBudget variable
	compBudget, err := strconv.Atoi(jobDetails.Budget)
	if err != nil {
		panic("error while conversion budget data from posting")
	}
	// Check if applicantBudget is greater than compBudget
	if applicantBudget > compBudget {
		fmt.Println("failed in budget")
		return false, models.JobApplicantResponse{}

	}
	// convert string to integer and store in compMinNoticePeriod variable
	compMinNoticePeriod, err := strconv.Atoi(jobDetails.MinNoticePeriod)
	fmt.Println(compMinNoticePeriod)
	if err != nil {
		panic("error while conversion min notice  period data from hr posting")
	}
	// convert string to integer and store in compMaxNoticePeriod variable
	compMaxNoticePeriod, err := strconv.Atoi(jobDetails.MaxNoticePeriod)
	fmt.Println(compMaxNoticePeriod)
	if err != nil {
		panic("error while conversion max notice period data from hr posting")
	}
	fmt.Println(validateApplication.Jobs.NoticePeriod)
	// convert string to integer and store in applicantNoticePeriod variable
	applicantNoticePeriod, err := strconv.Atoi(validateApplication.Jobs.NoticePeriod)
	fmt.Println(applicantNoticePeriod)
	if err != nil {
		panic("error while conversion notice period from applicant")
	}
	// Check if applicantNoticePeriod is less than compMinNoticePeriod and greater than compMaxNoticePeriod
	if (applicantNoticePeriod < compMinNoticePeriod) || (applicantNoticePeriod > compMaxNoticePeriod) {
		fmt.Println("failed in notice")

		return false, models.JobApplicantResponse{}
	}
	// Check if validateApplication.Jobs.JobDescription is equal to jobDetails.JobDescription
	if validateApplication.Jobs.JobDescription != jobDetails.JobDescription {
		fmt.Println("failed in descrpitoim")

		return false, models.JobApplicantResponse{}
	}
	// headCounter := 0
	count := 0
	fmt.Println(validateApplication.Jobs.JobLocation)
	for _, v1 := range validateApplication.Jobs.JobLocation {
		count = 0
		for _, v2 := range jobDetails.JobLocation {
			fmt.Println(v2.Name)
			if v1 == v2.ID {
				count++
				// headCounter++
			}
		}
	}
	if count == 0 {
		fmt.Println("failed location")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.JobType {
		count = 0
		for _, v2 := range jobDetails.JobType {
			if v1 == v2.ID {
				count++
				// headCounter++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed Jobtype")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Qualification {
		count = 0
		for _, v2 := range jobDetails.Qualification {
			if v1 == v2.ID {
				count++
				// headCounter++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed qualification")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Shift {
		count = 0
		for _, v2 := range jobDetails.Shift {
			if v1 == v2.ID {
				count++
				// headCounter++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed shift")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Technology {
		count = 0
		for _, v2 := range jobDetails.Technology {
			if v1 == v2.ID {
				count++
				// headCounter++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed technology")
		return false, models.JobApplicantResponse{}
	}
	count = 0
	for _, v1 := range validateApplication.Jobs.WorkMode {
		count = 0
		for _, v2 := range jobDetails.WorkMode {
			if v1 == v2.ID {
				count++
				// headCounter++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed workmode")
		return false, models.JobApplicantResponse{}
	}

	return true, validateApplication
}

// if (validateApplication == models.JobApplicantResponse{}) {
// 	log.Error().Err(errors.New("no candidates meet requirments"))
// 	return false, models.JobApplicantResponse{}
// }
