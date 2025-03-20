package database

import (
	"github.com/alexandrealfa/products-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	fakePassword := "passtest"

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Alfa", "Alfa@mail.com", fakePassword)
	userDB := NewUser(db)

	if err = userDB.Create(user); err != nil {
		t.Error(err)
	}

	var userValidate *entity.User

	if err = db.First(&userValidate, "id = ?", user.Id).Error; err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, user.Name, userValidate.Name)
	assert.Equal(t, user.Id, userValidate.Id)
	assert.Equal(t, user.Email, userValidate.Email)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Alfa", "Alfa@mail.com", "passtest")
	userDB := NewUser(db)

	if er := userDB.Create(user); er != nil {
		t.Error(er)
	}

	userValidate, userError := userDB.findByEmail("Alfa@mail.com")

	if userError != nil {
		t.Error(userError)
	}

	assert.Nil(t, err)
	assert.Equal(t, user.Name, userValidate.Name)
}
