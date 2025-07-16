package app

import (
	"context"
	"fmt"
	MongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"ocpp-smart-charging/internal/config"
	"ocpp-smart-charging/internal/routes"
	"time"
)

type App struct {
	MongoDB    *MongoDriver.Database
	HttpServer *http.Server
}

// Run запускает приложение
func Run() *App {
	httpServer := RunHttpServer()
	mongoDB := ConnectToMongo()

	return &App{
		HttpServer: httpServer,
		MongoDB:    mongoDB,
	}
}

func RunHttpServer() *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppConfig.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	fmt.Printf("Server launched at port: %d\n", config.AppConfig.Port)

	routes.RegisterRoutes()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Error during http server launch:", err)
		}
	}()

	return server
}

func ConnectToMongo() *MongoDriver.Database {
	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.DBConnectionTimeout)
	defer cancel()

	client, err := MongoDriver.Connect(ctx, options.Client().ApplyURI(config.AppConfig.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Double-check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Access to required Database
	db := client.Database(config.AppConfig.MongoDatabase)

	log.Println("Connection to MongoDB established!")

	//loadBalancerID := primitive.NewObjectID()
	//loadBalancer, err := mongoActions.GetLoadBalancer(db, loadBalancerID)
	//if err != nil {
	//	log.Fatalf("Can't get LoadBalancer #%s: %+v", loadBalancerID, err)
	//}
	//log.Println("loadBalancer", loadBalancer)

	//result, err := mongo.CreateLoadBalancer(db)
	//if err != nil {
	//	log.Fatalf("Can't create LoadBalancer! %+v", err)
	//}
	//log.Printf("CreateLoadBalancer result: %+v", result)

	return db
}
