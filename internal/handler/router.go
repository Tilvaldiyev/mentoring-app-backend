package handler

import (
	_ "github.com/Tilvaldiyev/mentoring-app-backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	go h.listenToWsChannel()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/countries", h.getCountries)
	router.GET("/languages", h.getLanguages)
	router.GET("/expertises", h.getExpertises)
	router.GET("/user-types", h.getUserTypes)
	router.GET("/levels", h.getUserLevels)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/forgot-password", h.sendSecretCode)
		auth.POST("/verify-secret-code", h.verifySecretCode)
		auth.PUT("/recover", h.recoverPassword)
	}

	api := router.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			article := apiV1.Group("/article")
			{
				article.GET("/:id", h.getArticle)
				article.POST("/create", h.createArticle)
				article.PUT("/:id", h.updateArticle)
				article.DELETE("/:id", h.deleteArticle)
			}

			post := apiV1.Group("/post")
			{
				post.GET("/:id", h.getPost)
				post.POST("/create", h.createPost)
				post.DELETE("/:id", h.deletePost)
			}

			mentor := apiV1.Group("/mentors")
			{
				mentor.GET("", h.getMentorsList)
			}

			ws := apiV1.Group("/ws")
			{
				ws.GET("", h.wsHandler)
			}
		}
	}

	return router
}
