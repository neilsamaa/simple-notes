package routes

import (
	"notes/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/notes", controllers.CreateNotesController)
	e.PUT("/notes/:id", controllers.UpdateNotesController)
	e.GET("/notes", controllers.GetNotesController)
	e.DELETE("/notes/:id", controllers.DeleteNotesController)
}
