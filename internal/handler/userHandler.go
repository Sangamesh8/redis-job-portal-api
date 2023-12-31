package handler

import (
	"encoding/json"
	"net/http"

	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *Handler) LoginUser(c *gin.Context) {
	// Extract trace ID from the request context
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		 // Log an error and respond with a 500 Internal Server Error if trace ID is missing
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	// Decode JSON request body into a NewUser struct
	var userData models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email and password",
		})
		return
	}
	// Perform user sign-in using the service
	token, err := h.service.UserSignIn(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	  // Respond with a JSON containing the generated token on successful user sign-in
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (h *Handler) SignUpUser(c *gin.Context) {
	ctx := c.Request.Context()

	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var userData models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email and password",
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email and password",
		})
		return
	}

	userDetails, err := h.service.UserSignup(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userDetails)

}

func (h *Handler) ForgotPassword(c *gin.Context) {
	// Extract trace ID from the request context
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		// Log an error and respond with a 500 Internal Server Error if trace ID is missing
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Decode JSON request body into a ForgotPasswordRequest struct
	var forgotPasswordRequest models.ForgotPasswordRequest
	err := json.NewDecoder(c.Request.Body).Decode(&forgotPasswordRequest)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email",
		})
		return
	}

	// Perform forgot password operation using the service
	err = h.service.ForgotPassword(ctx, forgotPasswordRequest)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Password reset link has been sent to your email",
	})
}

func (h* Handler) PasswordRecovery(c *gin.Context){
	// Extract trace ID from the request context
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		// Log an error and respond with a 500 Internal Server Error if trace ID is missing
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var passwordRecoveryRequest models.PasswordRecoveryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&passwordRecoveryRequest)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid reset token and new password",
		})
		return
	}
	// Perform password recovery operation using the service
	err = h.service.PasswordRecovery(ctx, passwordRecoveryRequest)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Password reset successful",
	})
}