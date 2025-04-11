package routes

import (
	"recu/PERSONA/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas con Gin
func SetupRoutes(router *gin.Engine) {


	product := router.Group("/personas")
	// Rutas para productos
	product.POST("/", controllers.SavePerson)
	product.GET("/OJITO", controllers.GetPersons)
	product.GET("/totalH",controllers.GetPersonsH)
	product.GET("/totalM",controllers.GetPersonsM)
}
