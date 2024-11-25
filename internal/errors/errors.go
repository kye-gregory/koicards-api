package errors

import (
	"log"

	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

// Fields
var (
	usernameField = *errpkg.NewErrorField("username")
	passwordField = *errpkg.NewErrorField("password")
	emailField = *errpkg.NewErrorField("email")
)

func plainError (code, message string) errpkg.StructuredError {
	return *errpkg.NewStructuredError(code, message)
}

func fieldError (code, message string, field errpkg.ErrorField) errpkg.StructuredError {
	return *errpkg.NewStructuredError(code, message).WithField(field)
}

// Internal Error
func Internal(errStack errpkg.ErrorStack, err error) {
	log.Print(err.Error())
	structuredErr := plainError("Internal", "internal error")
	errStack.InternalError(structuredErr)
}

// Username Errors
func UsernameEmpty(message string) errpkg.StructuredError { return fieldError("username_empty", message, usernameField) }
func UsernameLength(message string) errpkg.StructuredError { return fieldError("username_invalid_length", message, usernameField) }
func UsernameCharset(message string) errpkg.StructuredError { return fieldError("username_invalid_charset", message, usernameField) }
func UsernameInUse(message string) errpkg.StructuredError { return fieldError("username_in_use", message, usernameField) }

// Password Errors
func PasswordEmpty(message string) errpkg.StructuredError { return fieldError("password_empty", message, passwordField) }
func PasswordLength(message string) errpkg.StructuredError { return fieldError("password_invalid_length", message, passwordField) }

// Email Errors
func EmailInvalid(message string) errpkg.StructuredError { return fieldError("email_invalid", message, emailField) }
func EmailInUse(message string) errpkg.StructuredError { return fieldError("email_in_use", message, emailField) }

// Auth Errors
func AuthInvalidToken(message string) errpkg.StructuredError { return plainError("auth_invalid_token", message) }
