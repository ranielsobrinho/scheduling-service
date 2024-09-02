package models

import "github.com/google/uuid"

type SchedulingModel struct {
	Id           uuid.UUID `json:"id"`
	ScheduleDate string    `json:"schedule_date"`
	User         string    `json:"user_id"`
	Service      string    `json:"service"`
}
