// Paquete que contiene los controladores (lógica de negocio) de la tienda
package controllers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"larry-store/config"
	"larry-store/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Muestra la página principal con la lista de productos
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("products")           // Accede a la colección de productos
	cursor, err := collection.Find(context.TODO(), bson.D{}) // Busca todos los productos
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}

	var products []models.Product // Variable para guardar los productos
	if err = cursor.All(context.TODO(), &products); err != nil {
		http.Error(w, "Error al procesar productos", http.StatusInternalServerError)
		return
	}

	// Renderiza la plantilla HTML con los productos
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, products)
}

// Muestra la página de compra de un producto
func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("product") // Obtiene el nombre del producto desde la URL
	data := struct{ ProductName string }{ProductName: productName}
	// Renderiza la plantilla de checkout
	tmpl := template.Must(template.ParseFiles("templates/checkout.html"))
	tmpl.Execute(w, data)
}

// Procesa el formulario de pedido y guarda el pedido en la base de datos
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm() // Parsea los datos del formulario
	if err != nil {
		http.Error(w, "Error en el formulario", http.StatusBadRequest)
		return
	}

	// Crea un nuevo pedido con los datos del formulario
	newOrder := models.Order{
		CustomerName: r.FormValue("customerName"),
		Address:      r.FormValue("address"),
		ProductName:  r.FormValue("productName"),
		Status:       "Pendiente",
		Date:         time.Now(),
	}

	// Guarda el pedido en la colección de pedidos
	collection := config.DB.Collection("orders")
	_, err = collection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error al guardar pedido", http.StatusInternalServerError)
		return
	}

	// Renderiza la plantilla de éxito
	tmpl := template.Must(template.ParseFiles("templates/success.html"))
	tmpl.Execute(w, nil)
}
