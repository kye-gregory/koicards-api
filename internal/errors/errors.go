package errors

import (
	"log"

	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

func plainError (code errpkg.ErrorCode, message string) errpkg.StructuredError {
	return *errpkg.NewStructuredError(code, message)
}

func fieldError (code errpkg.ErrorCode, message string, field errpkg.ErrorField) errpkg.StructuredError {
	return *errpkg.NewStructuredError(code, message).WithField(field)
}

// Internal Error
var InternalCode = *errpkg.NewErrorCode("internal")
func Internal(errStack errpkg.ErrorStack, err error) {
	log.Print(err.Error())
	structuredErr := plainError(InternalCode, "internal error")
	errStack.InternalError(structuredErr)
}

// Username Errors
var usernameField = *errpkg.NewErrorField("username")
var UsernameLengthCode = *errpkg.NewErrorCode("username_invalid_length")
var UsernameCharsetCode = *errpkg.NewErrorCode("username_invalid_charset")
var UsernameInUseCode = *errpkg.NewErrorCode("username_in_use")
func UsernameLength(message string) errpkg.StructuredError { return fieldError(UsernameLengthCode, message, usernameField) }
func UsernameCharset(message string) errpkg.StructuredError { return fieldError(UsernameCharsetCode, message, usernameField) }
func UsernameInUse(message string) errpkg.StructuredError { return fieldError(UsernameInUseCode, message, usernameField) }

// Password Errors
var passwordField = *errpkg.NewErrorField("password")
var PasswordLengthCode = *errpkg.NewErrorCode("password_invalid_length")
func PasswordLength(message string) errpkg.StructuredError { return fieldError(PasswordLengthCode, message, passwordField) }

// Email Errors
var emailField = *errpkg.NewErrorField("email")
var EmailInvalidCode = *errpkg.NewErrorCode("email_invalid")
var EmailInUseCode = *errpkg.NewErrorCode("email_in_use")
func EmailInvalid(message string) errpkg.StructuredError { return fieldError(EmailInvalidCode, message, emailField) }
func EmailInUse(message string) errpkg.StructuredError { return fieldError(EmailInUseCode, message, emailField) }

// Auth Errors
var AuthInvalidTokenCode = *errpkg.NewErrorCode("auth_invalid_token")
func AuthInvalidToken(message string) errpkg.StructuredError { return plainError(AuthInvalidTokenCode, message) }