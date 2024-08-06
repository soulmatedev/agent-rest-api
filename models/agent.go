package models

type Agent struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}