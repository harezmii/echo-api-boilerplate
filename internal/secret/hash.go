package secret

import (
	configGetter "api/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), configGetter.Get("PASSWORD_COST", "int").(int))
	if err != nil {
		return ""
	}
	return string(generateFromPassword)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
