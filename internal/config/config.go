package config

import (
	"time"
)

// Config - config structure
type Config struct {
	MongoURI            string        // MongoDB connection URI
	MongoUser           string        // MongoDB user
	MongoPassword       string        // MongoDB user password
	DBConnectionTimeout time.Duration // DB requests timeout
	RequestTimeout      time.Duration // DB requests timeout
}

// AppConfig - config example
var AppConfig = Config{
	MongoURI:            "mongodb://localhost:27017",
	MongoUser:           "operator",
	MongoPassword:       "QWERTY",
	DBConnectionTimeout: 5 * time.Second,
	RequestTimeout:      5 * time.Second,
}
