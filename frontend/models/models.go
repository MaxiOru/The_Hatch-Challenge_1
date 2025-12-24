package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	Image       string             `bson:"image"`
}

type Order struct {
	CustomerName string    `bson:"customerName"`
	Address      string    `bson:"address"`
	ProductName  string    `bson:"productName"`
	Status       string    `bson:"status"`
	Date         time.Time `bson:"date"`
}
