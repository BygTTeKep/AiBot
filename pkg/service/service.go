package service

import "github.com/g91TeJl/AiBot/pkg/repository"

type Data interface {
	CollectData(username string, chatid int64, message string, answer []string) error
	GetNumberOfUsers() (int64, error)
}

type Service struct {
	Data
}

func NewService(repo *repository.Repositroy) *Service {
	return &Service{
		Data: NewDataService(repo),
	}
}
