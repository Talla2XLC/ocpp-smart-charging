package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoadBalancer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // Unique ID
	Name     string             `bson:"name"`          // Power load balancer name
	Tenant   string             `bson:"tenant"`        // Owner tenant
	LimitKW  float64            `bson:"limit_kw"`      // Max available power in kW
	LimitAmp float64            `bson:"limit_kw"`      // Max available current in Amp
	Active   bool               `bson:"active"`        // Is balancer active
}
