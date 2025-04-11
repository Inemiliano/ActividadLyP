package application

import "recu/PERSONA/domain"

type GetTotalH struct {
	repo domain.Iperson
}

func NewGetTotalH(repo domain.Iperson) *GetTotalH {
	return &GetTotalH{repo: repo}
}

func (c *GetTotalH) Execute() (int,error) {
	return c.repo.GetCountH()
}