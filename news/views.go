package news

import (
	"github.com/gin-gonic/gin"
)

// ListView ...
func ListView(c *gin.Context) {
	n := News{}
	c.JSON(200, gin.H{
		"status": true,
		"news":   n.Objects().All(),
	})
}
