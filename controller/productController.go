package controller

import (
	"database/sql"
	"ecommerce/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InitDb(database *sql.DB) {
	db = database
}

func ReadAllProduct(c *gin.Context) {
	log.Println("Fetching all products...")
	stmt := "select * from product"
	rows, err := db.Query(stmt)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Amount, &p.Qty, &p.Image, &p.IDCategory)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	// Deklarasi variable untuk menyimpan data request
	var product model.Product

	// Bind JSON dari request body ke variable product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Query untuk menyimpan data ke database
	query := `INSERT INTO product (name, description, amount, qty, image, id_category) 
			  VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, product.Name, product.Description, product.Amount, product.Qty, product.Image, product.IDCategory)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Ambil ID dari product yang baru disimpan
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product ID"})
		return
	}

	product.ID = id // Set ID ke object product

	// Kirim response
	c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Validasi ID (harus berupa angka)
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid product ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Query untuk menghapus produk berdasarkan ID
	query := "DELETE FROM product WHERE id = ?"
	result, err := db.Exec(query, productID)
	if err != nil {
		log.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	// Cek apakah ada baris yang dihapus
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func UpdateProduct(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Validasi ID (harus berupa angka)
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid product ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Bind JSON dari request body ke struct Product
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Query untuk memperbarui data produk
	query := `UPDATE product 
			  SET name = ?, description = ?, amount = ?, qty = ?, image = ?, id_category = ? 
			  WHERE id = ?`
	result, err := db.Exec(query, product.Name, product.Description, product.Amount, product.Qty, product.Image, product.IDCategory, productID)
	if err != nil {
		log.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Cek apakah ada baris yang diperbarui
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
