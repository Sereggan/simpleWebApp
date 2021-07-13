package store

import "github.com/Sereggan/simpleWebApp/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
