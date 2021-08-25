package service

type User interface {
	Login(username string, password string) error
}
