package models

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	AgentID     int    `json:"agent_id"`
}