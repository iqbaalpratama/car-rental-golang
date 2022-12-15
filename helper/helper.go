package helper

import (
	"car-rental/model"
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func GenerateToken(data model.Token) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = data.ID
	claims["role"] = data.Role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenData(c *gin.Context) (model.Token, error) {
	var data model.Token
	claims := jwt.MapClaims{}
	tokenString := ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return model.Token{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		data = model.Token{}
		role := fmt.Sprint(claims["role"])
		uid, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return model.Token{}, err
		}
		data.ID = int(uid)
		data.Role = role
		return data, nil
	}
	return model.Token{}, nil
}
