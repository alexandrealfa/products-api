package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrealfa/products-api/internal/database"
	"github.com/alexandrealfa/products-api/internal/dto"
	"github.com/alexandrealfa/products-api/internal/entity"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type UserHandler struct {
	userDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		userDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
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

func (H *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := H.userDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := H.Jwt.Encode(map[string]interface{}{
		"sub":             u.Id.String(),
		"expiration_date": time.Now().Add(time.Second * time.Duration(H.JwtExpiresIn)).Unix(),
	})
	fmt.Println("AQUII", tokenString)
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(accessToken)
}
