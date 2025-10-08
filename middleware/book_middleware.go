package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/Learn_Gin_Framework/models"
)

func BookMiddleware(c *gin.Context) {
	if len(models.Books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no books yet"})
		c.Abort()
		return
	}
	c.Next()
}
