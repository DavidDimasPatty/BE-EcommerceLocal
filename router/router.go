package router

import (
	"database/sql"
	"ecommerce/controller"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// SetupRouter mengatur routing untuk aplikasi
func SetupRouter() *gin.Engine {
	// Buat instance router
	r := gin.Default()

	// Koneksi database (hardcoded untuk contoh)
	db, err := sql.Open("mysql", "root:!@tcp(localhost:3306)/ecommerce")
	if err != nil {
		log.Println(err)
		panic("Failed to connect to database")
	}

	// Inisialisasi database di controller
	controller.InitDb(db)

	// Define routes
	r.GET("/products", controller.ReadAllProduct)
	r.POST("/products", controller.CreateProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)
	r.PUT("/products/:id", controller.UpdateProduct)
	return r
}
