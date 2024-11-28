package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	pass  = "1234"
	name  = "test_user"
	email = "test@tst.com"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(name, email, pass)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.password)
	assert.True(t, user.ValidatePassword(pass))
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser(name, email, pass)

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pass))
	assert.False(t, user.ValidatePassword(pass+"1"))
	assert.NotEqual(t, pass+"1", user.password)
}
