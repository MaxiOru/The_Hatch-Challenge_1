package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Estructura para los productos de la tienda
// Colección en MongoDB: "products"
// Nota: En Go, el nombre de la colección se especifica en el controlador al consultar la base de datos.
type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	Image       string             `bson:"image"`
}

// Estructura para los pedidos realizados por los clientes
// Colección en MongoDB: "orders"
// Nota: En Go, el nombre de la colección se especifica en el controlador al consultar la base de datos.
type Order struct {
	CustomerName string    `bson:"customerName"`
	Address      string    `bson:"address"`
	ProductName  string    `bson:"productName"`
	Status       string    `bson:"status"`
	Date         time.Time `bson:"date"`
}
