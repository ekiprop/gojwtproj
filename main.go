package main

import (
	"log"
	"os"

	"github.com/ekiprop/gojwtproj/handlers"
	"github.com/ekiprop/gojwtproj/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)

	router := r.Group("/api")

	router.POST("/register", server.Register)
	router.POST("/login", server.Login)

	return r

}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := SetupRouter()

	log.Fatal(r.Run(":" + port))

}
