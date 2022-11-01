package applicationscontroller

import "github.com/gin-gonic/gin"

func GetDetectableApplications(c *gin.Context) {
	c.JSON(200, []any{})
}
