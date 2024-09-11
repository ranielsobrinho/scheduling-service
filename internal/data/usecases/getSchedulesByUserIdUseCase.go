package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type GetSchedulesByUserIdUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewGetSchedulesByUserIdUseCase(scheduleRepository repositories.SchedulingRepository) GetSchedulesByUserIdUseCase {
	return GetSchedulesByUserIdUseCase{scheduleRepository: scheduleRepository}
}

func (getSchedulesByUserIdUseCase *GetSchedulesByUserIdUseCase) GetSchedulesByUserId(userId int) ([]models.SchedulingModel, error) {
	return getSchedulesByUserIdUseCase.scheduleRepository.GetSchedulesByUserId(userId)
}
