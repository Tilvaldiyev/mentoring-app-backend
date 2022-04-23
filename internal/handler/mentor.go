package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	expertiseColumn = "expertise_id"
	countryColumn   = "country_id"
	languageColumn  = "language_id"
	searchTerm      = "search_term"
)

// @Summary Get mentors list
// @Tags Mentor
// @Security ApiKeyAuth
// @Description Get mentors list
// @Accept json
// @Produce json
// @Param expertise query string  false  "search by expertise"
// @Param country query string  false  "search by country"
// @Param language query string  false  "search by language"
// @Success 200 {array} model.MentorResponse
// @Success 204
// @Failure 400,422 {object} model.Response
// @Router /api/v1/mentors [GET]
func (h *Handler) getMentorsList(c *gin.Context) {
	searchMap := make(map[string]string, 3)
	expertise, isExist := c.GetQuery("expertise")
	if isExist {
		searchMap[expertiseColumn] = expertise
	}

	country, isExist := c.GetQuery("country")
	if isExist {
		searchMap[countryColumn] = country
	}

	lang, isExist := c.GetQuery("language")
	if isExist {
		searchMap[languageColumn] = lang
	}

	term, isExist := c.GetQuery("term")
	if isExist {
		searchMap[searchTerm] = term
	}

	mentors, err := h.services.GetMentorsList(searchMap)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	if len(mentors) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, mentors)
}
