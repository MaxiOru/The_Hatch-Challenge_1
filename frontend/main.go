package main

import (
	"log"
	"net/http"

	"larry-store/config"
	"larry-store/controllers"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDB()

	// Servir archivos estáticos (imágenes)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rutas principales de la tienda
	http.HandleFunc("/", controllers.IndexHandler)            // Página principal (lista de productos)
	http.HandleFunc("/checkout", controllers.CheckoutHandler) // Página de compra de producto
	http.HandleFunc("/order", controllers.OrderHandler)       // Procesa el pedido

	port := "8080" // Puerto

	// Iniciar el servidor web en el puerto especificado
	log.Printf("Tienda corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
