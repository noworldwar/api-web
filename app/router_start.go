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

	// player
	r.GET("/", ctr.Index)
	r.GET("/login", ctr.Login)
	r.GET("/logout", ctr.Logout)
	r.GET("/register", ctr.Register)
	r.POST("/login", ctr.LoginRequest)
	r.POST("/register", ctr.RegisterRequest)

	// game
	r.GET("/pgsoft", ctr.PGSoft)
	r.GET("/livegame", ctr.LiveGame)
	r.GET("/wglobby", ctr.WGLobby)

	// wallet
	r.POST("/validate", ctr.Validate)
	r.POST("/balance", ctr.WalletBalance)
	r.POST("/debit", ctr.WalletDebit)
	r.POST("/credit", ctr.WalletCredit)
	r.POST("/rollback", ctr.Rollback)

	model.WGServer = http.Server{Addr: ":80", Handler: r}
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
