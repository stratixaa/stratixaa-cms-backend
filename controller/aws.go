package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	awssettings "yadhronics-blog/aws-settings"
	"yadhronics-blog/response"
	"yadhronics-blog/settings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func AwsPresignedURL(c *fiber.Ctx) error {
	allowedExtensions := map[string]bool{
		"png":  true,
		"jpg":  true,
		"jpeg": true,
		"svg":  true,
	}

	fileExt := c.Query("ext") // or from filename
	if !allowedExtensions[fileExt] {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Invalid file type. Only png, jpg, jpeg, svg allowed",
			ErrorTime:    time.Now(),
		})
	}

	contentTypeMap := map[string]string{
		"png":  "image/png",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"svg":  "image/svg+xml",
	}

	contentType := contentTypeMap[fileExt]

	objectKey := fmt.Sprintf("object-%d.%s", time.Now().UnixNano(), fileExt)

	presignedURL, err := awssettings.PreSignedURL.PresignPutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket:      aws.String(settings.Config.AWS.BucketName),
			Key:         aws.String(objectKey),
			ContentType: aws.String(contentType),
		},
		s3.WithPresignExpires(15*time.Minute),
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    fiber.StatusOK,
		StatusMessage: "success",
		Data: fiber.Map{
			"presigned_url": presignedURL.URL,
			"object_key":    settings.Config.AWS.CloudfrontDomain + "/" + objectKey,
		},
	})
}
