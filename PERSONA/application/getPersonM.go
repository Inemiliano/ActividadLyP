package application

import "recu/PERSONA/domain"

type GetTotalM struct {
	repo domain.Iperson
}

func NewGetTotalM(repo domain.Iperson) *GetTotalM {
	return &GetTotalM{repo: repo}
}

func (c *GetTotalM) Execute() (int,error) {
	return c.repo.GetCountM()
}