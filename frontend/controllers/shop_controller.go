package controllers

import (
	"context"
	"html/template"
	"net/http"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"larry-store/config"
	"larry-store/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("products")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}

	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		http.Error(w, "Error al procesar productos", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, products)
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("product")
	data := struct{ ProductName string }{ProductName: productName}
	tmpl := template.Must(template.ParseFiles("templates/checkout.html"))
	tmpl.Execute(w, data)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error en el formulario", http.StatusBadRequest)
		return
	}

	newOrder := models.Order{
		CustomerName: r.FormValue("customerName"),
		Address:      r.FormValue("address"),
		ProductName:  r.FormValue("productName"),
		Status:       "Pendiente",
		Date:         time.Now(),
	}

	collection := config.DB.Collection("orders")
	_, err = collection.InsertOne(context.TODO(), newOrder)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error al guardar pedido", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/success.html"))
	tmpl.Execute(w, nil)
}
