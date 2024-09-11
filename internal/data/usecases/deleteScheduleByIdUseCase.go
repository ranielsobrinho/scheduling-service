package usecases

import "github.com/ranielsobrinho/scheduling-service-api/internal/infra/database/repositories"

type DeleteScheduleByIdUseCase struct {
	scheduleRepository repositories.SchedulingRepository
}

func NewDeleteScheduleByIdUseCase(scheduleRepository repositories.SchedulingRepository) DeleteScheduleByIdUseCase {
	return DeleteScheduleByIdUseCase{scheduleRepository: scheduleRepository}
}

func (deleteScheduleByIdUseCase *DeleteScheduleByIdUseCase) DeleteScheduleById(scheduleId string) error {
	err := deleteScheduleByIdUseCase.scheduleRepository.DeleteScheduleById(scheduleId)

	if err != nil {
		return err
	}

	return nil
}
