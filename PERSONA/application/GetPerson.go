package application

import "recu/PERSONA/domain"

type GetAllPerson struct {
	repo domain.Iperson
}

func NewGetAll(repo domain.Iperson)*GetAllPerson{
	return &GetAllPerson{repo: repo}
}

func (c *GetAllPerson)Execute()([]domain.Person,error){
	return c.repo.GetAll()
}