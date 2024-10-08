package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ranielsobrinho/scheduling-service-api/internal/domain/models"
)

type SchedulingRepository struct {
	connection *sql.DB
}

func NewSchedulingRepository(connection *sql.DB) SchedulingRepository {
	return SchedulingRepository{connection: connection}
}

func (schedulingRepository *SchedulingRepository) GetSchedules() ([]models.SchedulingModel, error) {
	query := "SELECT * FROM seucarlos.schedules"

	rows, err := schedulingRepository.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []models.SchedulingModel{}, err
	}

	var schedulingList []models.SchedulingModel
	var schedulingObj models.SchedulingModel

	for rows.Next() {
		err := rows.Scan(
			&schedulingObj.Id,
			&schedulingObj.ScheduleDate,
			&schedulingObj.User,
			&schedulingObj.Service,
			&schedulingObj.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.SchedulingModel{}, err
		}

		schedulingList = append(schedulingList, schedulingObj)
	}

	rows.Close()

	return schedulingList, nil
}

func (schedulingRepository *SchedulingRepository) CreateSchedule(schedule models.SchedulingModel) (string, error) {
	query, err := schedulingRepository.connection.Prepare("INSERT INTO seucarlos.schedules (id, schedule_date, user_id, service) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var schedulingId string
	var id = uuid.New()

	err = query.QueryRow(id, schedule.ScheduleDate, schedule.User, schedule.Service).Scan(&schedulingId)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()

	return schedulingId, nil
}

func (schedulingRepository *SchedulingRepository) GetSchedulesByUserId(userId int) ([]models.SchedulingModel, error) {
	query := "SELECT schedules.id, schedules.schedule_date, schedules.service, schedules.user_id, schedules.created_at from seucarlos.schedules WHERE user_id = $1"

	rows, err := schedulingRepository.connection.Query(query, userId)
	if err != nil {
		return []models.SchedulingModel{}, err
	}

	var schedulingList []models.SchedulingModel
	var schedulingObj models.SchedulingModel

	for rows.Next() {
		err := rows.Scan(
			&schedulingObj.Id,
			&schedulingObj.ScheduleDate,
			&schedulingObj.Service,
			&schedulingObj.User,
			&schedulingObj.CreatedAt,
		)

		if err != nil {
			return []models.SchedulingModel{}, err
		}

		schedulingList = append(schedulingList, schedulingObj)
	}

	rows.Close()

	return schedulingList, nil
}

func (schedulingRepository *SchedulingRepository) DeleteScheduleById(scheduleId string) error {
	query, err := schedulingRepository.connection.Prepare("DELETE FROM seucarlos.schedules WHERE id = $1")

	if err != nil {
		return err
	}

	query.QueryRow(scheduleId)

	query.Close()

	return nil
}

func (schedulingRepository *SchedulingRepository) GetScheduleById(scheduleId string) (models.SchedulingModel, error) {
	query, err := schedulingRepository.connection.Prepare("SELECT id, schedule_date, service, user_id, created_at FROM seucarlos.schedules WHERE id = $1")
	if err != nil {
		return models.SchedulingModel{}, err
	}

	var schedulingObj models.SchedulingModel

	err = query.QueryRow(scheduleId).Scan(&schedulingObj.Id, &schedulingObj.ScheduleDate, &schedulingObj.Service, &schedulingObj.User, &schedulingObj.CreatedAt)
	if err != nil {
		return models.SchedulingModel{}, err
	}

	query.Close()
	return schedulingObj, nil
}

func (SchedulingRepository *SchedulingRepository) UpdateScheduleById(scheduleId string, schedule models.SchedulingModel) (models.SchedulingModel, error) {
	query, err := SchedulingRepository.connection.Prepare("UPDATE seucarlos.schedules SET schedule_date = $1, service = $2 WHERE id = $3 RETURNING id, schedule_date, service, user_id, created_at")
	if err != nil {
		return models.SchedulingModel{}, err
	}

	var schedulingObj models.SchedulingModel

	err = query.QueryRow(schedule.ScheduleDate, schedule.Service, scheduleId).Scan(&schedulingObj.Id, &schedulingObj.ScheduleDate, &schedulingObj.Service, &schedulingObj.User, &schedulingObj.CreatedAt)
	if err != nil {
		return models.SchedulingModel{}, err
	}

	query.Close()
	return schedulingObj, nil
}

func (schedulingRepository *SchedulingRepository) GetSchedulesByDayDate(dayDate string) ([]models.SchedulingModel, error) {
	initialDate := dayDate + "T00:00:00"
	endDate := dayDate + "T23:59:59"
	query := "SELECT id, schedule_date, service, user_id, created_at FROM seucarlos.schedules WHERE schedule_date BETWEEN $1 AND $2"
	rows, err := schedulingRepository.connection.Query(query, initialDate, endDate)
	if err != nil {
		return []models.SchedulingModel{}, err
	}

	var schedulingList []models.SchedulingModel
	var schedulingObj models.SchedulingModel

	for rows.Next() {
		err := rows.Scan(
			&schedulingObj.Id,
			&schedulingObj.ScheduleDate,
			&schedulingObj.Service,
			&schedulingObj.User,
			&schedulingObj.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.SchedulingModel{}, err
		}

		schedulingList = append(schedulingList, schedulingObj)
	}

	rows.Close()

	return schedulingList, nil
}

func (schedulingRepository *SchedulingRepository) GetSchedulesByDayMonth(month string) ([]models.SchedulingModel, error) {
	query := "SELECT id, schedule_date, service, user_id, created_at FROM seucarlos.schedules WHERE EXTRACT(MONTH FROM schedule_date) = $1"
	rows, err := schedulingRepository.connection.Query(query, month)
	if err != nil {
		return []models.SchedulingModel{}, err
	}

	var schedulingList []models.SchedulingModel
	var schedulingObj models.SchedulingModel

	for rows.Next() {
		err := rows.Scan(
			&schedulingObj.Id,
			&schedulingObj.ScheduleDate,
			&schedulingObj.Service,
			&schedulingObj.User,
			&schedulingObj.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.SchedulingModel{}, err
		}

		schedulingList = append(schedulingList, schedulingObj)
	}

	rows.Close()

	return schedulingList, nil
}
