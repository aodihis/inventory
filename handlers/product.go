package handlers

import (
	"database/sql"
	"encoding/json"
	"inventory/models"
	"net/http"
)

func Create(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	_, err := db.Exec("INSERT INTO products (name, quantity, price) VALUES ($1, $2, $3)",
		product.Name, product.Qty, product.Price)

	if err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Product(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var product models.Product

	productId := r.FormValue("id")

	err := db.QueryRow("select id, name, quantity, price, created_at, updated_at from products where id=?", productId).Scan(&product.ID, &product.Name, &product.Qty, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		http.Error(w, "Cannot find product", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

func Products(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	row, err := db.Query("select id, name, quantity, price, created_at, updated_at from products")

	if err != nil {
		http.Error(w, "Failed to query products", http.StatusInternalServerError)
		return
	}

	for row.Next() {
		var product models.Product
		if err := row.Scan(&product.ID, &product.Name, &product.Qty, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			http.Error(w, "Failed to get products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(products)
}
