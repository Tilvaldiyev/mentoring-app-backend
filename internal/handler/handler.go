package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get list of countries
// @Tags
// @Description Get list of countries
// @Accept json
// @Produce json
// @Success 200 {array} model.Country
// @Failure 422 {object} model.Response
// @Router /countries [GET]
func (h *Handler) getCountries(c *gin.Context) {
	countries, err := h.services.GetCountries()
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, countries)
}

// @Summary Get list of languages
// @Tags
// @Description Get list of languages
// @Accept json
// @Produce json
// @Success 200 {array} model.Language
// @Failure 422 {object} model.Response
// @Router /languages [GET]
func (h *Handler) getLanguages(c *gin.Context) {
	countries, err := h.services.GetLanguages()
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, countries)
}

// @Summary Get list of expertises
// @Tags
// @Description Get list of expertises
// @Accept json
// @Produce json
// @Success 200 {array} model.Expertise
// @Failure 422 {object} model.Response
// @Router /expertises [GET]
func (h *Handler) getExpertises(c *gin.Context) {
	countries, err := h.services.GetExpertises()
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, countries)
}

// @Summary Get list of user types
// @Tags
// @Description Get list of user types
// @Accept json
// @Produce json
// @Success 200 {array} model.UserType
// @Failure 422 {object} model.Response
// @Router /user-types [GET]
func (h *Handler) getUserTypes(c *gin.Context) {
	userTypes, err := h.services.GetUserTypes()
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userTypes)
}

// @Summary Get list of user levels
// @Tags
// @Description Get list of user levels
// @Accept json
// @Produce json
// @Success 200 {array} model.Level
// @Failure 422 {object} model.Response
// @Router /expertises [GET]
func (h *Handler) getUserLevels(c *gin.Context) {
	userLevels, err := h.services.GetUserLevels()
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userLevels)
}
