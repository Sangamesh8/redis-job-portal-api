package service

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_ProcessJobApplication(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData []models.JobApplicantResponse
	}
	tests := []struct {
		name    string
		args    args
		want    []models.JobApplicantResponse
		wantErr bool
		setup   func(mockRepo *repository.MockUserRepo)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				jobData: []models.JobApplicantResponse{
					{
						Name:  "Sheladon",
						JobID: uint(15),
						Jobs: models.JobApplicant{
							CompanyID:      uint(1),
							Name:           "Go Developer",
							NoticePeriod:   "30",
							Budget:         "550000",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{2, 5},
							WorkMode:       []uint{1, 2},
							JobDescription: "gRPC",
							Qualification:  []uint{1},
							Shift:          []uint{1},
							JobType:        []uint{1, 2},
						},
					},
				},
			},
			want: []models.JobApplicantResponse{
				{
					Name:  "Sheladon",
					JobID: uint(15),
					Jobs: models.JobApplicant{
						CompanyID:      uint(1),
						Name:           "Go Developer",
						NoticePeriod:   "30",
						Budget:         "550000",
						JobLocation:    []uint{1, 2},
						Technology:     []uint{2, 5},
						WorkMode:       []uint{1, 2},
						JobDescription: "gRPC",
						Qualification:  []uint{1},
						Shift:          []uint{1},
						JobType:        []uint{1, 2},
					},
				},
			},
			wantErr: false,
			setup: func(mockRepo *repository.MockUserRepo) {
				mockRepo.EXPECT().JobByID(gomock.Any(), uint(15)).Return(models.Jobs{
					Model: gorm.Model{ID: 15},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					CompanyID:       1,
					Name:            "Go Developer",
					MinNoticePeriod: "0",
					MaxNoticePeriod: "60",
					Budget:          "550000",
					JobLocation: []models.JobLocation{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
					Technology: []models.Technology{
						{Model: gorm.Model{ID: 2}},
						{Model: gorm.Model{ID: 5}},
					},
					WorkMode: []models.WorkMode{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
					JobDescription: "gRPC",
					Qualification: []models.Qualification{
						{Model: gorm.Model{ID: 1}},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: 1}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
				}, nil).Times(1)

			},
		},
		{
			name: "falied in checking notice period",
			args: args{
				ctx: context.Background(),
				jobData: []models.JobApplicantResponse{
					{
						Name:  "leonard",
						JobID: uint(15),
						Jobs: models.JobApplicant{
							CompanyID:      uint(1),
							Name:           "Go Developer",
							NoticePeriod:   "70",
							Budget:         "550000",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{2, 5},
							WorkMode:       []uint{1, 2},
							JobDescription: "gRPC",
							Qualification:  []uint{1},
							Shift:          []uint{1},
							JobType:        []uint{1, 2},
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			setup: func(mockRepo *repository.MockUserRepo) {
				mockRepo.EXPECT().JobByID(gomock.Any(), uint(15)).Return(models.Jobs{
					Model: gorm.Model{ID: 15},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					CompanyID:       1,
					Name:            "Go Developer",
					MinNoticePeriod: "0",
					MaxNoticePeriod: "60",
					Budget:          "550000",
					JobLocation: []models.JobLocation{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
					Technology: []models.Technology{
						{Model: gorm.Model{ID: 2}},
						{Model: gorm.Model{ID: 5}},
					},
					WorkMode: []models.WorkMode{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
					JobDescription: "gRPC",
					Qualification: []models.Qualification{
						{Model: gorm.Model{ID: 1}},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: 1}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: 1}},
						{Model: gorm.Model{ID: 2}},
					},
				}, nil).Times(1)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			tt.setup(mockRepo)
			s := &Service{
				UserRepo: mockRepo,
			}
			got, err := s.ProcessJobApplication(tt.args.ctx, tt.args.jobData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProcessJobApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ProcessJobApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx       context.Context
		jobData   models.CreateJobs
		CompanyID uint64
	}
	tests := []struct {
		name        string
		args        args
		want        models.ResponseForJobs
		wantErr     bool
		mockNewRepo func() (models.ResponseForJobs, error)
	}{
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				jobData: models.CreateJobs{
					CompanyID:       uint(1),
					Name:            "SDE 1",
					MinNoticePeriod: "0",
					MaxNoticePeriod: "60",
					Budget:          "600000",
					JobLocation: []uint{
						uint(1), uint(2),
					},
					Technology: []uint{
						uint(2), uint(5),
					},
					WorkMode: []uint{
						uint(1), uint(2),
					},
					JobDescription: "DSA",
					Qualification: []uint{
						uint(1),
					},
					Shift: []uint{
						uint(1), uint(2),
					},
				},
			},
			want: models.ResponseForJobs{},
			wantErr: true,
			mockNewRepo: func() (models.ResponseForJobs, error) {
				return models.ResponseForJobs{}, errors.New("test case error")
			},
		},
		{
			name: "success case",
			want: models.ResponseForJobs{},
			args: args{
				ctx: context.Background(),
				jobData: models.CreateJobs{
					CompanyID:       uint(1),
					Name:            "SDE 1",
					MinNoticePeriod: "0",
					MaxNoticePeriod: "60",
					Budget:          "600000",
					JobLocation: []uint{
						uint(1), uint(2),
					},
					Technology: []uint{
						uint(2), uint(5),
					},
					WorkMode: []uint{
						uint(1), uint(2),
					},
					JobDescription: "DSA",
					Qualification: []uint{
						uint(1),
					},
					Shift: []uint{
						uint(1), uint(2),
					},
					JobType: []uint{
						uint(1),
					},
				},
				},
				wantErr: false,
				mockNewRepo : func()(models.ResponseForJobs,error){
					return models.ResponseForJobs{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockNewRepo != nil {
				mockRepo.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockNewRepo()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.CompanyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
