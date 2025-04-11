package application

import "recu/PERSONA/domain"

type AddPerson struct {
	repo domain.Iperson
}


func NewAddPerson(repo domain.Iperson)*AddPerson{
	return &AddPerson{repo: repo}
}

func(c*AddPerson)Execute(persona *domain.Person)error{
	return c.repo.Save(persona)
}