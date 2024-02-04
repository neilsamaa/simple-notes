package controllers

import (
	"net/http"
	"notes/configs"
	"notes/models"

	"github.com/labstack/echo/v4"
)

func CreateNotesController(c echo.Context) error {
	var notesRequest models.Notes
	c.Bind(&notesRequest)

	result := configs.DB.Create(&notesRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
		Status:  true,
		Data:    notesRequest,
	})
}

func GetNotesController(c echo.Context) error {
	var notes []models.Notes
	result := configs.DB.Find(&notes)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    notes,
	})
}

func UpdateNotesController(c echo.Context) error {
	id := c.Param("id")
	var notesRequest models.Notes
	if err := c.Bind(&notesRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	var existingNotes models.Notes
	if err := configs.DB.First(&existingNotes, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	existingNotes.Title = notesRequest.Title
	existingNotes.Content = notesRequest.Content
	existingNotes.Owner = notesRequest.Owner

	result := configs.DB.Save(&existingNotes)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mengubah data",
		Status:  true,
		Data:    existingNotes,
	})
}

func DeleteNotesController(c echo.Context) error {
	id := c.Param("id")
	var notes models.Notes

	if err := configs.DB.First(&notes, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "catatan tidak ditemukan",
			Status:  false,
			Data:    nil,
		})
	}

	notes.IsDeleted = true
	result := configs.DB.Save(&notes)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menghapus data",
		Status:  true,
		Data:    notes,
	})
}
