package repositories

import (
	"agent-rest-api/models"
	"database/sql"
)

type AgentRepository struct {
	DB *sql.DB
}

func NewAgentRepository(db *sql.DB) *AgentRepository {
	return &AgentRepository{DB: db}
}

func (ar *AgentRepository) CreateAgent(agent *models.Agent) error {
	var id int
	err := ar.DB.QueryRow("INSERT INTO Agent (name, status) VALUES ($1, $2) RETURNING id", agent.Name, agent.Status).Scan(&id)
	if err != nil {
		return err
	}
	agent.ID = id
	return nil
}

func (ar *AgentRepository) UpdateAgent(agent *models.Agent) error {
	_, err := ar.DB.Exec("UPDATE Agent SET name=$1, status=$2 WHERE id=$3", agent.Name, agent.Status, agent.ID)
	return err
}

func (ar *AgentRepository) DeleteAgent(id int) error {
	_, err := ar.DB.Exec("DELETE FROM Agent WHERE id=$1", id)
	return err
}

func (ar *AgentRepository) GetAgentByID(id int) (*models.Agent, error) {
	var agent models.Agent
	err := ar.DB.QueryRow("SELECT id, name, status FROM Agent WHERE id=$1", id).Scan(&agent.ID, &agent.Name, &agent.Status)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (ar *AgentRepository) GetAllAgents() ([]models.Agent, error) {
	rows, err := ar.DB.Query("SELECT id, name, status FROM Agent")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []models.Agent
	for rows.Next() {
		var agent models.Agent
		if err := rows.Scan(&agent.ID, &agent.Name, &agent.Status); err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}
	return agents, nil
}