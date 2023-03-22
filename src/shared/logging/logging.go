package logging

import (
	"log"
	"src/modules/logs/data"
	"src/modules/logs/models"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		var err error
		var log_id string

		repo := data.NewLogData()

		_log := models.Log{
			UserAgent: c.Request.UserAgent(),
			IP:        c.Request.RemoteAddr,
			Path:      c.Request.URL.Path,
			Method:    c.Request.Method,
		}

		if log_id, err = repo.Add(&_log); err != nil {
			log.Print(err.Error())
			c.AbortWithStatus(500)
			return
		}

		c.Set("log_id", log_id)
		c.Next()

	}

}

func LogSetAppID(appID string, c *gin.Context) {

	id := LogGetID(c)
	var err error

	if id == "" {
		panic("empty id")
	}

	if err = data.NewLogData().SetAppID(appID, id); err != nil {
		panic(err)
	}

}

func LogSetUserID(userID string, c *gin.Context) {

	id := LogGetID(c)
	var err error

	if id == "" {
		panic("empty id")
	}

	if err = data.NewLogData().SetUserID(userID, id); err != nil {
		panic(err)
	}

}

func LogSetAccess(appID string, userID string, c *gin.Context) {

	id := LogGetID(c)
	var err error

	if id == "" {
		panic("empty id")
	}

	item := &models.Log{UserID: userID, AppID: appID}

	if err = data.NewLogData().SetAccess(item, id); err != nil {
		panic(err)
	}

}

func LogClose(status int, message string, c *gin.Context) {

	var err error

	id := LogGetID(c)

	if id == "" {
		panic("empty id")
	}

	item := &models.Log{
		Message: message,
		Status:  status,
	}

	if err = data.NewLogData().SetResult(item, id); err != nil {
		panic(err)
	}

}

func LogGetID(c *gin.Context) string {

	if id, ok := c.Get("log_id"); ok {
		return id.(string)
	}

	return ""
}
