package config

import (
	"time"
)

// Config - config structure
type Config struct {
	Port                int
	MongoURI            string        // MongoDB connection URI
	MongoUser           string        // MongoDB user
	MongoPassword       string        // MongoDB user password
	MongoDatabase       string        // MongoDB database name
	DBConnectionTimeout time.Duration // DB connection timeout
	RequestTimeout      time.Duration // DB requests timeout
}

// AppConfig - config example
var AppConfig = Config{
	Port:                8080,
	MongoURI:            "mongodb://localhost:27017",
	MongoUser:           "operator",
	MongoPassword:       "QWERTY",
	MongoDatabase:       "smart-charging",
	DBConnectionTimeout: 5 * time.Second,
	RequestTimeout:      5 * time.Second,
}
