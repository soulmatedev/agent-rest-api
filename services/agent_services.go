package services

import (
	"agent-rest-api/models"
	"agent-rest-api/repositories"
)

type AgentService struct {
	Repo *repositories.AgentRepository
}

func NewAgentService(repo *repositories.AgentRepository) *AgentService {
	return &AgentService{Repo: repo}
}

func (as *AgentService) CreateAgent(agent *models.Agent) error {
	return as.Repo.CreateAgent(agent)
}

func (as *AgentService) UpdateAgent(agent *models.Agent) error {
	return as.Repo.UpdateAgent(agent)
}

func (as *AgentService) DeleteAgent(id int) error {
	return as.Repo.DeleteAgent(id)
}

func (as *AgentService) GetAgentByID(id int) (*models.Agent, error) {
	return as.Repo.GetAgentByID(id)
}

func (as *AgentService) GetAllAgents() ([]models.Agent, error) {
	return as.Repo.GetAllAgents()
}