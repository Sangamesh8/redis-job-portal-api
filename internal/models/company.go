package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
	Type     string `json:"type" validate:"required"`
}
type Jobs struct {
	//Original Jobs struct
	gorm.Model
	Company         Company         `json:"-" gorm:"ForeignKey:company_id"`
	CompanyID       uint            `json:"company_id"`
	Name            string          `json:"title"`
	MinNoticePeriod string          `json:"min_notice_period"`
	MaxNoticePeriod string          `json:"max_notice_period"`
	Budget          string          `json:"budget"`
	JobLocation     []JobLocation   `gorm:"many2many:job_location;"` //gorm
	Technology      []Technology    `gorm:"many2many:technology;"`
	WorkMode        []WorkMode      `gorm:"many2many:workmode;"`
	JobDescription  string          `json : "jobdescription"`
	Qualification   []Qualification `gorm:"many2many:qualification;"`
	Shift           []Shift         `gorm:"many2many:shift;"`
	JobType         []JobType       `gorm:"many2many:jobtype;"`
}

type CreateJobs struct {
	//CreateJobs struct need to pass it to JObs for storing it into database
	CompanyID       uint   `json:"company_id"`
	Name            string `json:"title"`
	MinNoticePeriod string `json:"min_notice_period"`
	MaxNoticePeriod string `json:"max_notice_period"`
	Budget          string `json:"budget"`
	JobLocation     []uint `json:"job_location"`
	Technology      []uint `json : "technology"`
	WorkMode        []uint `json : "workmode"`
	JobDescription  string `json : "jobdescription"`
	Qualification   []uint `json : "qualification"`
	Shift           []uint `json : "shift"`
	JobType         []uint `json : "jobtype"`
}

type JobApplicant struct {
	//json body struct for applicants to filter out the
	CompanyID      uint   `json:"company_id"`
	Name           string `json:"title"`
	NoticePeriod   string `json:"noticePeriod"`
	Budget         string `json:"budget"`
	JobLocation    []uint `json:"job_location"`
	Technology     []uint `json:"technology"`
	WorkMode       []uint `json:"workmode"`
	JobDescription string `json:"jobdescription"`
	Qualification  []uint `json:"qualification"`
	Shift          []uint `json:"shift"`
	JobType        []uint `json:"jobtype"`
}
type JobApplicantResponse struct {
	Name  string       `json:"name"`
	JobID uint         `json:"jobId"`
	Jobs  JobApplicant `json:"jobApplication"`
}
type ResponseForJobs struct {
	ID uint `json:"id"`
}
type JobLocation struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
type Technology struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type WorkMode struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
type Qualification struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Shift struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type JobType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
