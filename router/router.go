package router

import (
	"database/sql"
	"ecommerce/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// SetupRouter mengatur routing untuk aplikasi
func SetupRouter() *gin.Engine {
	// Buat instance router
	r := gin.Default()

	//.env read
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error Loading .env file : %v", errEnv)
	}

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Koneksi database (hardcoded untuk contoh)
	db, err := sql.Open("mysql", ""+dbUser+":"+dbPass+"@tcp("+dbHost+":"+port+")/ecommerce")
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
