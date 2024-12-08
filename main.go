package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Wearedead99!@tcp(localhost:3306)/ecommerce" // Ganti dengan username, password, dan nama database Anda
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Cek koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	// Menampilkan pesan jika berhasil terkoneksi
	fmt.Println("Successfully connected to MySQL!")
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	r.GET("/getProducts", func(c *gin.Context) {
		// Mengambil data produk dari database
		rows, err := db.Query("SELECT id, name FROM product") // Ganti dengan query yang sesuai
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error querying database",
			})
			return
		}
		defer rows.Close()

		var products []map[string]interface{}

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatal("Error scanning row: ", err)
			}

			// Menambahkan data ke slice products
			products = append(products, map[string]interface{}{
				"id":   id,
				"name": name,
			})
		}

		// Cek apakah ada error setelah iterasi
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error iterating over rows",
			})
			return
		}

		// Mengembalikan data produk sebagai response
		c.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	})

	r.Run(":8080")
}
