package controllers

import (
	"net/http"
	"recu/PERSONA/application"
	"recu/PERSONA/infraestructure"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPersonsM(c *gin.Context) {
	repo := infraestructure.NewMongoPersonRepository()
	useCase := application.NewGetTotalM(repo)

	timeout := time.After(30 * time.Second) // Tiempo máximo de espera (30 segundos)
	ticker := time.NewTicker(2 * time.Second) // Revisar cada 2 segundos

	defer ticker.Stop()

	// Obtener el total inicial de personas con sexo "M"
	initialTotal, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo ejecutar el método"})
		return
	}

	for {
		select {
		case <-timeout:
			// Si pasa el tiempo sin cambios, devolvemos el valor actual
			c.JSON(http.StatusOK, gin.H{"total": initialTotal})
			return
		case <-ticker.C:
			// Revisamos si el total ha cambiado
			currentTotal, err := useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo ejecutar el método"})
				return
			}

			if currentTotal != initialTotal {
				// Si hay un cambio, respondemos con el nuevo valor
				c.JSON(http.StatusOK, gin.H{"total": currentTotal})
				return
			}
		}
	}
}
