package util

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		return fallback
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GenerateUUID() string {
	return uuid.New().String()
}

func BcryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(id string, memberLevel uint) (string, error) {
	expiryTime := GetEnv("JWT_EXPIRY", "86400")

	token_lifespan, err := strconv.Atoi(expiryTime)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["iss"] = "auth.bilibili.com"
	claims["aud"] = "https://bilibili.com"
	claims["sub"] = id
	claims["nbf"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(token_lifespan)).Unix()

	var scope string

	// memberLevel => 0: admin, 1: seller, 2: buyer
	switch memberLevel {
	case 0:
		scope = "admin category"
	case 1:
		scope = "product invoice user"
	default:
		scope = "invoice user"
	}
	claims["scope"] = scope

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := GetEnv("JWT_KEY", "3ea7ac596cb207cb2ec4dffcabb1130bec74a08c")

	return token.SignedString([]byte(jwtSecret))
}

func VerifyJWT(token, scopeTarget string) error {
	jwtSecret := GetEnv("JWT_KEY", "3ea7ac596cb207cb2ec4dffcabb1130bec74a08c")

	tokenResult, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if _, err := matchTokenScope(*tokenResult, scopeTarget); err != nil {
		return err
	}

	return nil
}

func ExtractTokenId(token string) string {
	jwtSecret := GetEnv("JWT_KEY", "3ea7ac596cb207cb2ec4dffcabb1130bec74a08c")

	tokenResult, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	claims, _ := tokenResult.Claims.(jwt.MapClaims)

	return claims["sub"].(string)
}

func matchTokenScope(token jwt.Token, scopeTarget string) (bool, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	scopesList := strings.Split(claims["scope"].(string), " ")

	if !valueExists(scopesList, scopeTarget) {
		return false, errors.New("Unauthorized")
	}

	return true, nil
}

func valueExists(arr []string, value string) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
