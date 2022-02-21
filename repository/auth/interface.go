package auth

import "app_airbnb/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
