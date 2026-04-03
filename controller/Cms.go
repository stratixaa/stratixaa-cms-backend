package controller

import (
	"context"
	"net/http"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/response"
	"yadhronics-blog/service"

	"github.com/gofiber/fiber/v2"
)

func GetCMSData(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := c.Query("key", "home") // Default to "home" if no ID is provided

	//fetch data from DB
	result, err := service.GetCMSData(ctx, idParam)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          result,
	})
}

func UpdateCMSData(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := c.Query("key")

	//parsing a request body
	var cms map[string]interface{}
	if err := c.BodyParser(&cms); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Failed to parse request body",
			ErrorTime:    time.Now(),
		})
	}

	//saving data in db
	result, err := service.UpdateCMSData(ctx, cms, idParam)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response with the created objectid
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          result,
	})
}
