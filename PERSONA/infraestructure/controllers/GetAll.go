package controllers

import (
	"net/http"
	"recu/PERSONA/application"
	"recu/PERSONA/infraestructure"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context){

	repo:=infraestructure.NewMongoPersonRepository()
	useCase:=application.NewGetAll(repo)

	persona,err:=useCase.Execute();

	if err!=nil{
		c.JSON(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK,persona)

}