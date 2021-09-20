package store

import "github.com/vlasove/materials/tasks_2/utils/calendar/internal/app/models"

// Store ...
type Store struct {
	db         map[int]*models.Event
	repository *EventRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Open ...
func (s *Store) Open() error {
	db := make(map[int]*models.Event)
	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() error {
	return nil
}

// EventRepository ...
func (s *Store) EventRepository() *EventRepository {
	if s.repository == nil {
		s.repository = &EventRepository{store: s}
	}
	return s.repository
}
