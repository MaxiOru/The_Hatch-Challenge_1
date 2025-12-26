// Paquete para la configuración y conexión a la base de datos
package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variable global para acceder a la base de datos desde otros archivos
var DB *mongo.Database

func ConnectDB() {
	// Crear un contexto con timeout para la conexión a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Liberar recursos del contexto al finalizar

	// URI de conexión a MongoDB (local, hardcodeada)
	mongoURI := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err) // Si falla la conexión, termina el programa
	}

	// Seleccionar la base de datos (nombre hardcodeado)
	dbName := "larry_shop"
	DB = client.Database(dbName)
	log.Println("Conectado a MongoDB (Frontend Go)") // Mensaje de éxito
}
