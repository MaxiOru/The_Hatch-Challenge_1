package main

import (
	"log"
	"net/http"
	"os"

	"larry-store/config"
	"larry-store/controllers"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando valores por defecto")
	}

	config.ConnectDB()

	// Servir archivos estáticos (imágenes)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/checkout", controllers.CheckoutHandler)
	http.HandleFunc("/order", controllers.OrderHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Tienda corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
