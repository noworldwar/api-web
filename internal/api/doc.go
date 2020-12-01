package api

import (
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/api-web/internal/model"
)

func Index(c *gin.Context) {
	switch lang := c.Query("lang"); lang {
	case "cn":
		c.HTML(200, "index_cn.html", model.APIdocCn)
	case "en":
		c.HTML(200, "index_en.html", model.APIdocEn)
	default:
		c.HTML(200, "index_zh.html", model.APIdocZh)
	}
}

func DocIntegration(c *gin.Context) {
	switch lang := c.Query("lang"); lang {
	case "cn":
		c.HTML(200, "integration.html", model.APIdocCn)
	case "en":
		c.HTML(200, "integration.html", model.APIdocEn)
	default:
		c.HTML(200, "integration.html", model.APIdocZh)
	}
}

func DocHistory(c *gin.Context) {
	switch lang := c.Query("lang"); lang {
	case "cn":
		c.HTML(200, "dochistory_cn.html", model.APIdocCn)
	case "en":
		c.HTML(200, "dochistory_en.html", model.APIdocEn)
	default:
		c.HTML(200, "dochistory_zh.html", model.APIdocZh)
	}
}

func DocWorkflow(c *gin.Context) {
	switch lang := c.Query("lang"); lang {
	case "cn":
		c.HTML(200, "workflow_cn.html", model.APIdocCn)
	case "en":
		c.HTML(200, "workflow_en.html", model.APIdocEn)
	default:
		c.HTML(200, "workflow_zh.html", model.APIdocZh)
	}
}
