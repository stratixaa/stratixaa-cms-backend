package service

import (
	"context"
	"yadhronics-blog/database"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCMSData(ctx context.Context, key string) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := database.CMS.FindOne(ctx, bson.M{"key": key}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
