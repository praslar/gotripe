package handler

import (
	"github.com/gotripe/router/middleware"
	"github.com/gotripe/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	movie := v1.Group("/movie", jwtMiddleware)
	movie.POST("", h.CreateMovie)
	movie.GET("/all", h.FindAllMovies)
	movie.GET("/:movie_id", h.WatchMovie)

	stripe := v1.Group("/stripe")
	stripe.GET("/pubkey", h.GetPubKey)
	stripe.GET("/plan", h.GetPlan)
	// profiles := v1.Group("/profiles", jwtMiddleware)
	// profiles.GET("/:username", h.GetProfile)
	// profiles.POST("/:username/follow", h.Follow)
	// profiles.DELETE("/:username/follow", h.Unfollow)

	// articles := v1.Group("/articles", middleware.JWTWithConfig(
	// 	middleware.JWTConfig{
	// 		Skipper: func(c echo.Context) bool {
	// 			if c.Request().Method == "GET" && c.Path() != "/api/articles/feed" {
	// 				return true
	// 			}
	// 			return false
	// 		},
	// 		SigningKey: utils.JWTSecret,
	// 	},
	// ))
	// articles.POST("", h.CreateArticle)
	// articles.GET("/feed", h.Feed)
	// articles.PUT("/:slug", h.UpdateArticle)
	// articles.DELETE("/:slug", h.DeleteArticle)
	// articles.POST("/:slug/comments", h.AddComment)
	// articles.DELETE("/:slug/comments/:id", h.DeleteComment)
	// articles.POST("/:slug/favorite", h.Favorite)
	// articles.DELETE("/:slug/favorite", h.Unfavorite)
	// articles.GET("", h.Articles)
	// articles.GET("/:slug", h.GetArticle)
	// articles.GET("/:slug/comments", h.GetComments)

	// tags := v1.Group("/tags")
	// tags.GET("", h.Tags)
}
