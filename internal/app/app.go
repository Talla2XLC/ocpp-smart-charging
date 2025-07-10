package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"ocpp-smart-charging/internal/config"
	mongoActions "ocpp-smart-charging/internal/database/mongo"
	"ocpp-smart-charging/internal/routes"
)

type App struct {
	MongoClient *mongo.Client
}

func NewApp() *App {
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.DBConnectionTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Double-check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB established!")

	// Получаем доступ к конкретной базе данных
	db := client.Database("smart_charging") // Укажите имя вашей базы данных

	// Пример получения зарядных станций по ID балансировщика
	loadBalancerID := primitive.NewObjectID() // Замените на фактический ID
	loadBalancer, err := mongoActions.GetLoadBalancer(db, loadBalancerID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("loadBalancer", loadBalancer)

	return &App{
		MongoClient: client,
	}
}

// Run запускает приложение
func (app *App) Run() {
	routes.RegisterRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error during http server launch:", err)
	}
}
