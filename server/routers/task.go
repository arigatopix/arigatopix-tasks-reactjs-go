package routers

import (
	"net/http"
	"strconv"

	s "server/services"

	"github.com/gin-gonic/gin"
)

type AddTaskForm struct {
	Text     string `form:"text" json:"text" binding:"required" validate:"min=1,max=500"`
	Day      string `form:"day" json:"day" binding:"required"`
	Reminder bool   `form:"reminder" json:"reminder" validate:"boolean"`
}

// @Summary Create Task
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/ [post]
func AddTask(ctx *gin.Context) {
	var form *AddTaskForm

	err := ctx.ShouldBindJSON(&form)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{
		Text:     form.Text,
		Day:      form.Day,
		Reminder: form.Reminder,
	}

	task, err := todoService.Add()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})

}

// @Summary Get task with ID
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [get]
func GetTask(ctx *gin.Context) {

	// converse string to int
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	// create instance
	todoService := s.Task{ID: id}

	// call service with instance
	task, err := todoService.Get()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// @Summary Get multiple Tasks
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks [get]
func GetTasks(ctx *gin.Context) {
	todoService := s.Task{}

	tasks, err := todoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tasks,
	})
}

// @Summary Update Task by id
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [put]
func UpdateTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var form *AddTaskForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{
		ID:       id,
		Text:     form.Text,
		Day:      form.Day,
		Reminder: form.Reminder,
	}

	task, err := todoService.Update()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    task,
	})
}

// @Summary Delete Task by id
// @Produce  json
// @Success 200
// @Failure 400
// @Router /api/tasks/:id [delete]
func DeleteTask(ctx *gin.Context) {
	// converse string to int
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	todoService := s.Task{ID: id}

	if err := todoService.Delete(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
