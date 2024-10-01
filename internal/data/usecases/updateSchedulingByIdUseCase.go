package usecases

import (
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
	"github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"
)

type UpdateSchedulingByIdUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewUpdateSchedulingByIdUseCase(scheduleRepository repositories.SchedulingRepository) UpdateSchedulingByIdUseCase {
	return UpdateSchedulingByIdUseCase{scheduleRepository: scheduleRepository}
}

func (updateSchedulingByIdUseCase *UpdateSchedulingByIdUseCase) UpdateScheduleById(schedulingId string, schedule models.SchedulingModel) (models.SchedulingModel, error) {
	updatedSchedule, err := updateSchedulingByIdUseCase.scheduleRepository.UpdateScheduleById(schedulingId, schedule)

	if err != nil {
		return models.SchedulingModel{}, err
	}

	return updatedSchedule, nil
}
