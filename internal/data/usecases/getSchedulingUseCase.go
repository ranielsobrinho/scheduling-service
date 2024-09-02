package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type GetSchedulesUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewGetSchedulesUseCase(scheduleRepository repositories.SchedulingRepository) GetSchedulesUseCase {
	return GetSchedulesUseCase{scheduleRepository: scheduleRepository}
}

func (getSchedulingUseCase *GetSchedulesUseCase) GetSchedules() ([]models.SchedulingModel, error) {
	return getSchedulingUseCase.scheduleRepository.GetSchedules()
}
