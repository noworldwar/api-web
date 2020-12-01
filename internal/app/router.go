package app

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	api "github.com/noworldwar/api-web/internal/api"
	"github.com/noworldwar/api-web/internal/model"
)

func InitRouter() {
	r := gin.Default()
	r.Use(api.CheckToken())

	// 頁面
	r.HTMLRender = loadTemplates()

	// @Api:Wallet@
	r.POST("/validate", api.Validate)
	r.POST("/balance", api.WalletBalance)
	r.POST("/debit", api.WalletDebit)
	r.POST("/credit", api.WalletCredit)
	r.POST("/rollback", api.Rollback)
	// @end

	model.WGServer = http.Server{Addr: ":7531", Handler: r}
	r.Static("/static", "doc/static")
	r.Static("/test", "test")
	r.HTMLRender = loadTemplates()

	r.GET("/doc", api.Index)
	r.GET("/doc/integration", api.DocIntegration)
	r.GET("/doc/history", api.DocHistory)
	r.GET("/doc/workflow", api.DocWorkflow)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad Request"})
	})
}

func RunRouter() {
	if err := model.WGServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	includes, err := filepath.Glob("doc/page/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), "doc/layout/base.html", include)
	}

	return r
}
