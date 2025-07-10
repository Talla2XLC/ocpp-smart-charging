package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChargingStation struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`    // Unique ID
	LoadBalancerID primitive.ObjectID `bson:"load_balancer_id"` // To what LoadBalancer belongs this item
	EVSEID         string             `bson:"evse_id"`          // EVSE ID
	ConnectorID    string             `bson:"connector_id"`     // Connector ID
	Priority       int                `bson:"priority"`         // Current item Priority
	MaxKW          float64            `bson:"max_kw"`           // Max power that can be delivered
	MaxAmp         float64            `bson:"max_kw"`           // Max current that can be delivered
	MinLimitKW     float64            `bson:"min_limit_kw"`     // Min power that must be delivered
	MinLimitAmp    float64            `bson:"min_limit_kw"`     // Min current that must be delivered
	ActiveLimitKW  float64            `bson:"active_limit_kw"`  // Active power limit
	ActiveLimitAmp float64            `bson:"active_limit_amp"` // Active current limit
}
