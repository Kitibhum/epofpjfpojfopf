package service
import "github.com/GDSC-KMUTT/totp-session/repository"

type userService struct {
	repository repository.UserRepository // interface หรือ "Port" ของ repository
}

func NewUserService(userRepository repository.UserRepository) userService { // รับ "Adapter" ของ Repository
	return userService{repository: userRepository}
}

func (s userService) SignUp(email string, password string) (*string, *string, error) {
	// implement me
	s.repository.CreateUser(email, password, "")
	return nil, nil, nil
}

func (s userService) SignIn(email string, password string) (*User, error) {
	// implement me
	s.repository.CheckUser(email)
	return nil, nil
}
