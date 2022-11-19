package handler

import (
	"net/http"

	"github.com/GDSC-KMUTT/totp-session/service"
)

type userHandler struct {
	service service.UserService // interface หรือ "Port" จาก Service
}

func NewUserHandler(userService service.UserService) userHandler { // รับ "Adapter" ของ Service
	return userHandler{service: userService}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignUp("", "")
	w.Write([]byte("User created!"))
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignIn("", "")
	w.Write([]byte("User signed in!"))
}
