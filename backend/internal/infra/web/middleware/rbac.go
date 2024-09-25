package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("RBAC called: %s", time.Now())
		// c.Set("", "")
		c.Next()
	}
}
