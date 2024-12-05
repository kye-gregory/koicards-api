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
	"github.com/kye-gregory/koicards-api/pkg/auth"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type AuthService struct {
	store store.SessionStore
}


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
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil { errs.Internal(httpStack, err); return }

	// Setup Email
	to := []string{claims["email"].(string)}
	var body bytes.Buffer
	t, err := template.ParseFiles("internal/mail/templates/verify_email.html")
	if err != nil { errs.Internal(httpStack, err); return }

	// Create Template & Send
	t.Execute(&body, struct{VerificationLink string; Username string}{VerificationLink: "localhost:8080/api/v1/account/verify?token=" + signedToken, Username: username})
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
		return jwtSecretKey, nil
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


func (svc *AuthService) CreateSession(userID int, httpStack *errpkg.HttpStack) *models.Session {
	// Generate CSRF Token
	csrfToken, err := auth.GenerateCSRFToken()
	if err != nil { errs.Internal(httpStack, err); return nil }

	// Create And Store Session
	session := models.NewSession(*models.NewSessionData(userID, csrfToken))
	err = svc.store.CreateSession(session)
	if err != nil { errs.Internal(httpStack, err); return nil }

	return session
}


func (svc *AuthService) DeleteSession(sessionID string, httpStack *errpkg.HttpStack) {
	// Delete Session
	err := svc.store.DeleteSession(sessionID)
	if err != nil { errs.Internal(httpStack, err) }
}


func (svc *AuthService) VerifySession(sessionID string, csrfToken string, httpStack *errpkg.HttpStack) {
	// Get Session Data
	data, err := svc.store.GetSessionData(sessionID)
	if err != nil { errs.Internal(httpStack, err); return }
	
	// Compare CSRF Token
	structuredErr := errs.AuthUnauthorised("CSRF Token does not match")
	valid := data.CSRFToken == csrfToken
	if !valid { httpStack.Add(structuredErr); return }
}