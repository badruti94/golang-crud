package main

import (
	"golang-crud/controllers/item"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api/v1/notes", item.GetNotes)
	e.POST("/api/v1/notes", item.AddNote)
	e.PUT("/api/v1/notes/:id", item.UpdateNote)
	e.DELETE("/api/v1/notes/:id", item.DeleteNote)

	e.Logger.Fatal(e.Start(":3000"))

}
