package handler

import (
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.JobPortalService
}

func Api(a auth.Authentication, ser service.JobPortalService) *gin.Engine {
	r := gin.New()
	h := Handler{
		service: ser,
	}
	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
	}

	if err != nil {
		log.Panic("handlers not setup")
	}

	r.Use(m.Log(), gin.Recovery())

	r.GET("/check", m.Authenticate(Check))

	r.POST("/signup", h.SignUpUser)
	r.POST("/login", h.LoginUser)
	r.POST("/createCompany", m.Authenticate(h.CreateCompany))
	r.GET("/viewAllCompanies/all", m.Authenticate(h.ViewAllCompanies))
	r.GET("/viewCompanyByID/:id", m.Authenticate(h.ViewCompanyByID))
	r.GET("/job/viewByCompanyID/:cid", m.Authenticate(h.ViewJobByCompanyID))
	r.POST("createJobs/:cid", m.Authenticate(h.CreateJobs))
	r.GET("/viewJobs/all", m.Authenticate(h.ViewAllJobs))
	r.GET("/viewByJobID/:id", m.Authenticate(h.ViewJobByID))
	r.POST("/api/processapplication", m.Authenticate(h.ProcessJobApplication))
	return r
}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
