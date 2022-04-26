package item

import (
	"golang-crud/utils/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Note struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`
}

var dataResponse map[string]interface{}

func GetNotes(c echo.Context) error {
	var notes []Note

	db, err := db.Connect()
	if err != nil {
		return err
	}

	db.Find(&notes)

	dataResponse = map[string]interface{}{
		"data": map[string]interface{}{
			"notes": notes,
		},
	}

	return c.JSON(http.StatusOK, dataResponse)

}

func AddNote(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var note Note
	note.Title = c.FormValue("title")
	note.Note = c.FormValue("note")

	db.Select("Note", "Title").Create(&note)

	dataResponse = map[string]interface{}{
		"message": "Data saved",
	}

	return c.JSON(http.StatusOK, dataResponse)

}

func UpdateNote(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	var note Note

	note.Id, _ = strconv.Atoi(c.Param("id"))

	db.Model(&note).Updates(Note{
		Title: c.FormValue("title"),
		Note:  c.FormValue("note"),
	})

	dataResponse = map[string]interface{}{
		"message": "Data updated",
	}

	return c.JSON(http.StatusOK, dataResponse)

}

func DeleteNote(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	db.Delete(&Note{}, id)

	dataResponse = map[string]interface{}{
		"message": "Data deleted",
	}

	return c.JSON(http.StatusOK, dataResponse)

}
