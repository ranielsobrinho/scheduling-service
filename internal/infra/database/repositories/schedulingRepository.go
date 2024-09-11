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
	fmt.Println("Foi chamado aqui oia")
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
