package dbrepo

import "github.com/seemsod1/api-project/internal/models"

// AddSubscriber adds a new subscriber to the database
func (m *mockDB) AddSubscriber(subscriber models.Subscriber) error {
	args := m.Called(subscriber)
	return args.Error(0)
}

// GetSubscribers returns all subscribers from the database
func (m *mockDB) GetSubscribers(timezone int) ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
