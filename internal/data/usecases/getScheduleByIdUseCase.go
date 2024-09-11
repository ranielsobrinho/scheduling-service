package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type GetScheduleByIdUseCase struct {
	schedulingRepository repositories.SchedulingRepository
}

func NewGetScheduleByIdUseCase(schedulingRepository repositories.SchedulingRepository) GetScheduleByIdUseCase {
	return GetScheduleByIdUseCase{schedulingRepository: schedulingRepository}
}

func (getScheduleByIdUseCase *GetScheduleByIdUseCase) GetSchedulesById(scheduleId string) (models.SchedulingModel, error) {
	schedule, err := getScheduleByIdUseCase.schedulingRepository.GetScheduleById(scheduleId)
	if err != nil {
		return models.SchedulingModel{}, err
	}

	return schedule, nil
}
