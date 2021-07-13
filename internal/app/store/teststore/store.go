package teststore

import (
	"github.com/Sereggan/simpleWebApp/internal/app/model"
	_ "github.com/lib/pq"
)

type Store struct {
	userRepository *UserRepository
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
