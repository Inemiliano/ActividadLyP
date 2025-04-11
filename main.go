package main

import (
	"log"
	routsUser "recu/PERSONA/infraestructure/routes"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurar CORS para permitir cualquier origen
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // <-- Permitir todos los orígenes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false, // <-- true no es compatible con "*"
		MaxAge:           12 * time.Hour,
	}))

	routsUser.SetupRoutes(router)

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}
