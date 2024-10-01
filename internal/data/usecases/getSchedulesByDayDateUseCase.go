package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type GetSchedulesByDayDateUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewGetSchedulesByDayDateUseCase(scheduleRepository repositories.SchedulingRepository) GetSchedulesByDayDateUseCase {
	return GetSchedulesByDayDateUseCase{scheduleRepository: scheduleRepository}
}

func (getSchedulesByDayDateUseCase *GetSchedulesByDayDateUseCase) GetSchedulesByDayDate(dayDate string) ([]models.SchedulingModel, error) {
	schedules, err := getSchedulesByDayDateUseCase.scheduleRepository.GetSchedulesByDayDate(dayDate)

	if err != nil {
		return []models.SchedulingModel{}, err
	}

	return schedules, nil
}
