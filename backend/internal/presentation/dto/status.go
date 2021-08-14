package dto

import "time"

type Status struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Details string    `json:"details"`
	Time    time.Time `json:"time"`
}
