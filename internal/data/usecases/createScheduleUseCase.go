package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type CreateScheduleUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewCreateScheduleUseCase(scheduleRepository repositories.SchedulingRepository) CreateScheduleUseCase {
	return CreateScheduleUseCase{scheduleRepository: scheduleRepository}
}

func (createScheduleUseCase *CreateScheduleUseCase) CreateScheduleUseCase(schedule models.SchedulingModel) (string, error) {
	scheduleIdCreated, err := createScheduleUseCase.scheduleRepository.CreateSchedule(schedule)
	if err != nil {
		return "", err
	}

	return scheduleIdCreated, nil
}
