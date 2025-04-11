package controllers

import (
	"net/http"
	"recu/PERSONA/application"
	"recu/PERSONA/domain"
	"recu/PERSONA/infraestructure"

	"github.com/gin-gonic/gin"
)

func SavePerson(c *gin.Context){

	var Persona domain.Person

	if err:=c.ShouldBindJSON(&Persona);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err})
	}

	repo:=infraestructure.NewMongoPersonRepository()
	useCase:=application.NewAddPerson(repo)

	if err:=useCase.Execute(&Persona);err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":"no se creo a la persona"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"OK":"persona creado con exito"})

}