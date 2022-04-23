package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create article
// @Tags Article
// @Security ApiKeyAuth
// @Description Create article
// @Accept json
// @Produce json
// @Param input body model.Article true "request body"
// @Success 201
// @Failure 400,422 {object} model.Response
// @Router /api/v1/article [POST]
func (h *Handler) createArticle(c *gin.Context) {
	var req model.Article

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	if err := h.services.CreateArticle(req); err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Get article
// @Tags Article
// @Security ApiKeyAuth
// @Description Get article info by ID
// @Accept json
// @Produce json
// @Param id path string true "ID article"
// @Success 200 {object} model.Article
// @Failure 400,422 {object} model.Response
// @Router /api/v1/article/{id} [GET]
func (h *Handler) getArticle(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Errorf("incorrect id: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	article, err := h.services.GetArticle(idInt)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, article)

}

// @Summary Update article
// @Tags Article
// @Security ApiKeyAuth
// @Description Update article
// @Accept json
// @Produce json
// @Param id path string true "ID article"
// @Param input body model.UpdateArticleInput true "request body"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /api/v1/article/{id} [PUT]
func (h *Handler) updateArticle(c *gin.Context) {
	var req model.UpdateArticleInput

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Errorf("incorrect id: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	err = h.services.UpdateArticle(idInt, req)
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

// @Summary Delete article
// @Tags Article
// @Security ApiKeyAuth
// @Description Delete article
// @Accept json
// @Produce json
// @Param id path string true "ID article"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /api/v1/article/{id} [DELETE]
func (h *Handler) deleteArticle(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.log.Errorf("incorrect id: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	err = h.services.DeleteArticle(idInt)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Message:   "Deleted",
	})
}
