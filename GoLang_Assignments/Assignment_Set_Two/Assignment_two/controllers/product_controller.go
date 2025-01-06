package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"Assignment_two/models"
	"Assignment_two/utils"
	"github.com/gorilla/mux"
)

var db *sql.DB

func InitProductController(database *sql.DB) {
	db = database
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error adding product")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, "Product added successfully")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	query := `SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`
	row := db.QueryRow(query, id)

	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
		utils.RespondError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, product)
}

