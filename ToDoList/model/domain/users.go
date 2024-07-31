package domain

import "time"

type Users struct {
	Id         string
	Username   string
	Email      string
	Password   string
	CreatedAt time.Time
	UpdatedAt time.Time
}