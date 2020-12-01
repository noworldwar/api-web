package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckQueryData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.Query(v)) == "" {
			return v
		}
	}
	return ""
}

func IsInt(vals ...string) (b bool, s string) {
	for _, v := range vals {
		if v != "" {
			_, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				msg := v + " non-integer"
				return false, msg
			}
		}
	}

	return true, ""
}

func IsInt64(vals ...string) (b bool, s string) {
	for _, v := range vals {
		if v != "" {
			_, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				msg := v + " non-integer"
				return false, msg
			}
		}
	}

	return true, ""
}
