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

func UpdateCMSData(ctx context.Context, cms map[string]interface{}, key string) (map[string]interface{}, error) {
	// Update the document in the database
	_, err := database.CMS.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": cms})
	if err != nil {
		return nil, err
	}

	// Fetch the updated document
	var updatedData map[string]interface{}
	err = database.CMS.FindOne(ctx, bson.M{"key": key}).Decode(&updatedData)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
