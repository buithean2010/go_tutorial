package services

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() *loginService {
	return &loginService{
		authorizedUsername: "admin",
		authorizedPassword: "password",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
