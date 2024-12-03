package services

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/mail"
	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type AuthService struct {
	store store.SessionStore
}

// Constructor function for UserService
func NewAuthService(s store.SessionStore) *AuthService {
	return &AuthService{store: s}
}


func (svc *AuthService) SendEmailVerification(email string, username string, httpStack *errpkg.HttpStack) {
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
	if err != nil { errs.Internal(httpStack, err); return }

	// Setup Email
	to := []string{claims["email"].(string)}
	var body bytes.Buffer
	t, err := template.ParseFiles("internal/mail/templates/verify_email.html")
	if err != nil { errs.Internal(httpStack, err); return }

	// Create Template & Send
	t.Execute(&body, struct{VerificationLink string; Username string}{VerificationLink: "localhost:8080/api/v1/accounts/verify?token=" + signedToken, Username: username})
	err = mail.Send("KoiCards - Verify Email", body, to)
	if err != nil { errs.Internal(httpStack, err); return }
}


func (svc *AuthService) VerifyEmail(tokenString string, httpStack *errpkg.HttpStack) (string) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Check internal errors
	if err != nil { errs.Internal(httpStack, err); return "" }

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email
	}

	// Return invalid
	structuredErr := errs.AuthInvalidToken("invalid token")
	httpStack.Add(structuredErr)
	return ""
}


func (svc *AuthService) CreateSession(userID userVO.ID, httpStack *errpkg.HttpStack) *models.Session {
	// Create Session
	session, err := svc.store.CreateSession(userID)
	if err != nil { errs.Internal(httpStack, err); return nil }

	return session
}


func (svc *AuthService) DeleteSession(sessionID string, httpStack *errpkg.HttpStack) {
	// Delete Session
	err := svc.store.DeleteSession(sessionID)
	if err != nil { errs.Internal(httpStack, err) }
}


func (svc *AuthService) VerifySession(sessionID string, httpStack *errpkg.HttpStack) {
	structuredErr := errs.AuthUnauthorised("invalid session id")
	exists, err := svc.store.VerifySession(sessionID)
	if err != nil { errs.Internal(httpStack, err); return }
	if !exists { httpStack.Add(structuredErr) }
}