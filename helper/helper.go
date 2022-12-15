package helper

import (
	"net/mail"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsLengthPhoneNumberValid(phone string) bool {
	if len([]rune(phone)) < 11 || len([]rune(phone)) > 14 {
		return false
	}
	return true
}

func IsLengthPasswordValid(password string) bool {
	if len([]rune(password)) < 6 {
		return false
	}
	return true
}

func IsCarYearValid(year int) bool {
	if year < 2006 || year > 2022 {
		return false
	}
	return true
}

func IsUrlValid(imageUrl string) bool {
	u, err := url.Parse(imageUrl)
	return err == nil && u.Scheme != "" && u.Host != ""
}
