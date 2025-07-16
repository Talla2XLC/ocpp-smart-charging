package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"ocpp-smart-charging/internal/config"
	"ocpp-smart-charging/internal/models" // Импортируя модели
)

// CreateLoadBalancer create new LoadBalancer
func CreateLoadBalancer(db *mongo.Database) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.RequestTimeout) // Установите таймаут в 5 секунд
	defer cancel()

	result, err := db.Collection("load_balancers").InsertOne(ctx, bson.D{
		{"_id", primitive.NewObjectID()},
		{"active", true},
	})
	return result, err
}

// GetLoadBalancers get all LoadBalancers
func GetLoadBalancers(db *mongo.Database) ([]models.LoadBalancer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.RequestTimeout) // Установите таймаут в 5 секунд
	defer cancel()

	cursor, err := db.Collection("load_balancers").Find(ctx, bson.M{})
	defer func() {
		if closeErr := cursor.Close(ctx); closeErr != nil {
			log.Println("Failed to close cursor:", closeErr)
		}
	}()

	var loadBalancers []models.LoadBalancer

	// Decode cursor
	for cursor.Next(ctx) {
		var loadBalancer models.LoadBalancer
		if err := cursor.Decode(&loadBalancer); err != nil {
			return nil, err
		}
		loadBalancers = append(loadBalancers, loadBalancer)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loadBalancers, err
}

// GetLoadBalancer get LoadBalancer by ID
func GetLoadBalancer(db *mongo.Database, loadBalancerID primitive.ObjectID) (models.LoadBalancer, error) {
	var loadBalancer models.LoadBalancer

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.RequestTimeout) // Установите таймаут в 5 секунд
	defer cancel()

	err := db.Collection("load_balancers").FindOne(ctx, bson.D{{"_id", loadBalancerID}}).Decode(&loadBalancer)
	return loadBalancer, err
}

// GetLoadBalancerWithStations - Get LoadBalancer with ChargingStations by ID
func GetLoadBalancerWithStations(db *mongo.Database, loadBalancerID primitive.ObjectID) (models.LoadBalancer, []models.ChargingStation, error) {
	var loadBalancer models.LoadBalancer
	var chargingStations []models.ChargingStation

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.RequestTimeout) // Установите таймаут в 5 секунд
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", loadBalancerID}}}}, // Find required LoadBalancer
		{
			{"$lookup", bson.D{
				{"from", "charging_stations"},        // Collection name to join
				{"localField", "_id"},                // Field in LoadBalancer
				{"foreignField", "load_balancer_id"}, // Field in ChargingStation
				{"as", "chargingStations"},           // Alias for ChargingStations array
			}},
		},
	}

	collection, err := db.Collection("load_balancers").Aggregate(ctx, pipeline)
	if err != nil {
		return loadBalancer, chargingStations, err
	}
	defer func() {
		if err := collection.Close(ctx); err != nil {
			log.Println("Ошибка при закрытии курсора:", err)
		}
	}()

	// Извлекаем результат
	if collection.Next(ctx) {
		if err := collection.Decode(&loadBalancer); err != nil {
			return loadBalancer, chargingStations, err
		}

		// As we used $lookup, chargingStations will be a field inside loadBalancer
		if err := collection.Decode(bson.M{"stations": &chargingStations}); err != nil {
			return loadBalancer, chargingStations, err
		}
	}

	return loadBalancer, chargingStations, nil
}
