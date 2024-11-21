package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kye-gregory/koicards-api/internal/mail"
	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

var secretKey = []byte("your-secret-key") // TODO: Put this as env variables!

type AuthService struct {
}

// Constructor function for UserService
func NewAuthService() *AuthService {
	return &AuthService{}
}


func (svc *AuthService) SendEmailVerification(email string, username string, errStack *errorstack.HttpStack) {
	// Define claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),                     
		"type":  "email_verification",
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil { errStack.ReturnInternalError(); return }

	// Setup Email
	to := []string{claims["email"].(string)}
	var body bytes.Buffer
	t, err := template.ParseFiles("internal/mail/templates/verify_email.html")
	if err != nil { errStack.ReturnInternalError(); return }

	// Create Template & Send
	t.Execute(&body, struct{VerificationLink string; Username string}{VerificationLink: "localhost:8080/api/v1/accounts/verify?token=" + signedToken, Username: username})
	err = mail.Send("KoiCards - Verify Email", body, to)
	if err != nil { errStack.ReturnInternalError(); return }
}


func (svc *AuthService) VerifyEmail(tokenString string, errStack *errorstack.HttpStack) (string) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil { errStack.ReturnInternalError(); return "" }

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email
	}

	// Return Invalid
	err = errors.New("invalid token")
	errStack.Add("verification", err)
	return ""
}