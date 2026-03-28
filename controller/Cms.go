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
