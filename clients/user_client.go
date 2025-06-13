package clients

import (
	"errors"
	"sync"

	"testcodex/models"
)

// UserClient simulates a user database client using an in-memory map.
type UserClient struct {
	mu     sync.Mutex
	users  map[int]*models.User
	nextID int
}

// NewUserClient creates a new in-memory user client.
func NewUserClient() *UserClient {
	return &UserClient{users: make(map[int]*models.User), nextID: 1}
}

// Create stores a new user and returns it.
func (c *UserClient) Create(u *models.User) *models.User {
	c.mu.Lock()
	defer c.mu.Unlock()
	u.ID = c.nextID
	c.nextID++
	c.users[u.ID] = u
	return u
}

// Get retrieves a user by ID.
func (c *UserClient) Get(id int) (*models.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	u, ok := c.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// Update modifies an existing user.
func (c *UserClient) Update(id int, u *models.User) (*models.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	u.ID = id
	c.users[id] = u
	return u, nil
}

// Delete removes a user by ID.
func (c *UserClient) Delete(id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(c.users, id)
	return nil
}

// List returns all stored users.
func (c *UserClient) List() []*models.User {
	c.mu.Lock()
	defer c.mu.Unlock()
	users := make([]*models.User, 0, len(c.users))
	for _, u := range c.users {
		users = append(users, u)
	}
	return users
}
