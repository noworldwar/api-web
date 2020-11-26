package app

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	ctr "github.com/noworldwar/api-web/controller"
	"github.com/noworldwar/api-web/model"
)

func InitRouter() {
	r := gin.Default()
	r.Use(ctr.CheckToken())

	// js 與 css
	r.Static("/static", "view/static")

	// 頁面
	r.HTMLRender = loadTemplates()

	// wallet
	r.POST("/validate", ctr.Validate)
	r.POST("/balance", ctr.WalletBalance)
	r.POST("/debit", ctr.WalletDebit)
	r.POST("/credit", ctr.WalletCredit)
	r.POST("/rollback", ctr.Rollback)

	model.WGServer = http.Server{Addr: ":7531", Handler: r}
}

func RunRouter() {
	if err := model.WGServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loadTemplates() multitemplate.Renderer {

	r := multitemplate.NewRenderer()

	includes, err := filepath.Glob("view/page/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), "view/layout/base.html", include)
	}

	return r
}
