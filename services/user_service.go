package services

import (
	"testcodex/clients"
	"testcodex/models"
)

// UserService provides user operations.
type UserService struct {
	client *clients.UserClient
}

// NewUserService creates a new user service.
func NewUserService(c *clients.UserClient) *UserService {
	return &UserService{client: c}
}

func (s *UserService) Create(u *models.User) *models.User {
	return s.client.Create(u)
}

func (s *UserService) Get(id int) (*models.User, error) {
	return s.client.Get(id)
}

func (s *UserService) Update(id int, u *models.User) (*models.User, error) {
	return s.client.Update(id, u)
}

func (s *UserService) Delete(id int) error {
	return s.client.Delete(id)
}

func (s *UserService) List() []*models.User {
	return s.client.List()
}
