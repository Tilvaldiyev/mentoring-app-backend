package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"net/http"
)

// @Summary User registration
// @Tags Auth
// @Description User registration
// @Accept json
// @Produce json
// @Param input body model.Users true "request body"
// @Success 201
// @Failure 400,422 {object} model.Response
// @Router /auth/sign-up [POST]
func (h *Handler) signUp(c *gin.Context) {
	var req model.Users

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	if err := h.services.CreateUser(req); err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)

}

// @Summary Authorization
// @Tags Auth
// @Description User authorization
// @Accept json
// @Produce json
// @Param input body model.SignInInput true "request body"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /auth/sign-in [POST]
func (h *Handler) signIn(c *gin.Context) {
	var req model.SignInInput

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	 token, err := h.services.GenerateToken(req)
	 if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Message:   token,
	})
}

// @Summary Send secret code to email
// @Tags Auth
// @Description Send secret code to email when forgot password
// @Accept json
// @Produce json
// @Param input body model.PasswordRecoveryInput true "request body"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /auth/forgot-password [POST]
func (h *Handler) sendSecretCode(c *gin.Context) {
	var req model.PasswordRecoveryInput

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	response, err := h.services.VerifyEmailAndSendCode(req)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Message:   response,
	})

}

// @Summary Verify sent secret code
// @Tags Auth
// @Description Verify sent secret code
// @Accept json
// @Produce json
// @Param input body model.PasswordRecoveryInput true "request body"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /auth/verify-secret-code [POST]
func (h *Handler) verifySecretCode(c *gin.Context) {
	var req model.PasswordRecoveryInput

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	err := h.services.VerifySecretCode(req)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Message:   "verified",
	})

}

// @Summary Update password
// @Tags Auth
// @Description Update password
// @Accept json
// @Produce json
// @Param input body model.SignInInput true "request body"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /auth/recover [PUT]
func (h *Handler) recoverPassword(c *gin.Context) {
	var req model.SignInInput
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	err := h.services.RecoverPassword(req)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Message:   "Updated",
	})
}
