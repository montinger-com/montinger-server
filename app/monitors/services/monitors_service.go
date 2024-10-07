package monitors_service

import (
	monitors_model "github.com/montinger-com/montinger-server/app/monitors/models"
	monitors_repository "github.com/montinger-com/montinger-server/app/monitors/repositories"
	"github.com/montinger-com/montinger-server/lib/db"
)

type MonitorsService struct {
	monitorsRepo *monitors_repository.MonitorsRepository
}

func NewMonitorsService() *MonitorsService {

	return &MonitorsService{
		monitorsRepo: monitors_repository.NewMonitorsRepository(db.MongoClient),
	}
}

func (s *MonitorsService) GetAll() ([]*monitors_model.Monitor, error) {
	return s.monitorsRepo.GetAll()
}
