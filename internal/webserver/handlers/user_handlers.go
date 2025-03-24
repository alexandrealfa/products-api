package handlers

import (
	"encoding/json"
	"github.com/alexandrealfa/products-api/internal/database"
	"github.com/alexandrealfa/products-api/internal/dto"
	"github.com/alexandrealfa/products-api/internal/entity"
	"net/http"
)

type UserHandler struct {
	userDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{userDB}
}

func (H *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var Content dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&Content); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(Content.Name, Content.Email, Content.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = H.userDB.Create(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}
