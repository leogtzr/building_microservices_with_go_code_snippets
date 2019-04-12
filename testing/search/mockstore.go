package search

import "github.com/stretchr/testify/mock"

// MockStore ...
type MockStore struct {
	mock.Mock
}

// Search ...
func (m *MockStore) Search(name string) []User {
	args := m.Mock.Called(name)
	return args.Get(0).([]User)
}
