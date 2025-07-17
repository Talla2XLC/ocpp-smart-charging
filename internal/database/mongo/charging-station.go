package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ocpp-smart-charging/internal/config"
)

// AddChargingStation add new ChargingStation params to existing LoadBalancer
func AddChargingStation(db *mongo.Database, loadBalancerID string) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.RequestTimeout) // Установите таймаут в 5 секунд
	defer cancel()

	loadBalancerIdObj, err := primitive.ObjectIDFromHex(loadBalancerID)
	if err != nil {
		return nil, err
	}

	result, err := db.Collection("charging_stations").InsertOne(ctx, bson.D{
		{"_id", primitive.NewObjectID()},
		{"load_balancer_id", loadBalancerIdObj},
	})
	return result, err
}
