package response

import (
	"src/shared/logging"

	"github.com/gin-gonic/gin"
)

func SendError(status int, message string, c *gin.Context) {
	logging.LogClose(status, message, c)
	c.AbortWithStatus(status)
}

func SendOK(c *gin.Context) {
	logging.LogClose(201, "", c)
	c.AbortWithStatus(201)
}

func SendID(id string, c *gin.Context) {
	logging.LogClose(200, "", c)
	c.IndentedJSON(200, gin.H{
		"id": id,
	})
}

func SendToken(token string, c *gin.Context) {
	logging.LogClose(200, "", c)
	c.IndentedJSON(200, gin.H{
		"token": token,
	})
}

func SendData(data any, c *gin.Context) {
	logging.LogClose(200, "", c)
	c.IndentedJSON(200, data)
}

func SendList(total int, items any, c *gin.Context) {
	logging.LogClose(200, "", c)
	c.IndentedJSON(200, gin.H{
		"total": total,
		"items": items,
	})
}
