package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestFake(t *testing.T) {
	user := &User{}
	assert.Empty(t, user)

	user.Email = fake.EmailAddress()
	user.Name = fake.MaleFirstName()
	user.Password = fake.Password(9, 9, true, true, true)
	assert.True(t, govalidator.IsEmail(user.Email))
	assert.NotEqual(t, "", user.Name)
	assert.NotEqual(t, "", user.Password)
}
