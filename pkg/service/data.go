package service

import "github.com/g91TeJl/AiBot/pkg/repository"

type DataService struct {
	repo *repository.Repositroy
}

func NewDataService(repo *repository.Repositroy) *DataService {
	return &DataService{repo: repo}
}

func (s *DataService) CollectData(username string, chatid int64, message string, answer []string) error {
	return s.repo.CollectData(username, chatid, message, answer)
}

func (s *DataService) GetNumberOfUsers() (int64, error) {
	return s.repo.GetNumberOfUsers()
}
