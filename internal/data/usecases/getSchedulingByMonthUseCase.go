package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type GetSchedulesByMonthUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewGetSchedulesByMonthUseCase(scheduleRepository repositories.SchedulingRepository) GetSchedulesByMonthUseCase {
	return GetSchedulesByMonthUseCase{scheduleRepository: scheduleRepository}
}

func (getSchedulesByMonthUseCase *GetSchedulesByMonthUseCase) GetSchedulesByMonth(month string) ([]models.SchedulingModel, error) {
	schedules, err := getSchedulesByMonthUseCase.scheduleRepository.GetSchedulesByDayMonth(month)

	if err != nil {
		return []models.SchedulingModel{}, err
	}

	return schedules, nil
}
