package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create post
// @Tags Post
// @Security ApiKeyAuth
// @Description Create post
// @Accept json
// @Produce json
// @Param input body model.Posts true "request body"
// @Success 201
// @Failure 400,422 {object} model.Response
// @Router /api/v1/post [POST]
func (h *Handler) createPost(c *gin.Context) {
	var req model.Posts

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("JSON binding err: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	if err := h.services.CreatePost(req); err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Get post
// @Tags Post
// @Security ApiKeyAuth
// @Description Get post info by ID
// @Accept json
// @Produce json
// @Param id path string true "ID article"
// @Success 200 {object} model.Posts
// @Failure 400,422 {object} model.Response
// @Router /api/v1/post/{id} [GET]
func (h *Handler) getPost(c *gin.Context) {
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

	post, err := h.services.GetPost(idInt)
	if err != nil {
		h.log.Errorf("service err: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)

}

// @Summary Delete post
// @Tags Post
// @Security ApiKeyAuth
// @Description Delete post
// @Accept json
// @Produce json
// @Param id path string true "ID article"
// @Success 200 {object} model.Response
// @Failure 400,422 {object} model.Response
// @Router /api/v1/post/{id} [DELETE]
func (h *Handler) deletePost(c *gin.Context) {
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

	err = h.services.DeletePost(idInt)
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
		Message:   "Deleted",
	})
}
