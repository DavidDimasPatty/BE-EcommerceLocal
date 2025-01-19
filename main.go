package main

import (
	"ecommerce/router"
	"log"
)

func main() {
	// Inisialisasi router
	r := router.SetupRouter()

	// Jalankan server pada port 8080
	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
