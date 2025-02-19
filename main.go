package main

import (
	"log"

	"faqs/config"
	"faqs/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Gagal membaca file .env, menggunakan variabel lingkungan")
	}

	// Koneksi Database
	config.ConnectDB()

	// Inisialisasi Gin
	router := gin.Default()

	// Konfigurasi CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ganti dengan domain yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Tambahkan Routes
	routes.FAQRoutes(router)

	// Ambil port dari environment atau gunakan default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server berjalan di port " + port)
	router.Run(":" + port)
}
