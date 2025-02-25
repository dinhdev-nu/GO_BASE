package service

import (
	"github.com/dinhdev-nu/GO_BASE/internal/repo"
)

type PongService struct{
	pr *repo.PongRepo
}

func NewPongService( ) *PongService {
	return &PongService{
		pr: repo.NewPongRepo(),
	}	
}

func (ps *PongService) GetInforPong() string {
	return ps.pr.GetPong()
}