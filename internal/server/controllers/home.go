package controllers

import (
	"net/http"

	"terralist/pkg/api"
	"terralist/pkg/builders"
	"terralist/pkg/webui"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// HomeController registers the endpoints for portal home
// Either the login page, if the user is not authenticated
// or the home page
type HomeController interface {
	api.RestController
}

// DefaultHomeController is a concrete implementation of
// HomeController
type DefaultHomeController struct {
	UIManager webui.Manager
}

func (c *DefaultHomeController) Paths() []string {
	return []string{
		"", // empty path
	}
}

func (c *DefaultHomeController) Subscribe(apis ...*gin.RouterGroup) {
	api := apis[0]

	withLayout := builders.
		NewSliceBuilder[string]().
		Add("layout.tpl")

	homeKey, _ := c.UIManager.Register(
		withLayout.
			Add("home.tpl").
			Build(),
	)

	api.GET("/", func(ctx *gin.Context) {
		if err := c.UIManager.Render(ctx.Writer, homeKey, &map[string]string{
			"Provider": "GitHub",
		}); err != nil {
			log.Debug().AnErr("Error", err).Msg("Cannot render home view")
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	})

	api.POST("/", func(ctx *gin.Context) {
		// Call login route to start authentication process
		ctx.String(http.StatusOK, "Not yet implemented.")
	})
}
