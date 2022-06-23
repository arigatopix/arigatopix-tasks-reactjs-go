package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(CORSMiddleware())

	apiGroups := r.Group("/api")

	task := apiGroups.Group("/tasks")
	{
		// GET /tasks
		task.GET("/", GetTasks)

		// GET /tasks/:id
		task.GET("/:id", GetTask)

		// POST /tasks
		task.POST("/", AddTask)

		// PUT /tasks/:id
		task.PUT("/:id", UpdateTask)

		// DELETE /tasks/:id
		task.DELETE("/:id", DeleteTask)
	}

	r.Run(":5000")

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
