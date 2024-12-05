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
	log.Printf("INTERNAL ERROR: %s", err.Error())
	structuredErr := plainError(InternalCode, "internal error")
	errStack.InternalError(structuredErr)
}

// Username Errors
var usernameField = *errpkg.NewErrorField("username")
var UsernameLengthCode = *errpkg.NewErrorCode("username_invalid_length")
var UsernameFormatCode = *errpkg.NewErrorCode("username_invalid_format")
var UsernameCharsetCode = *errpkg.NewErrorCode("username_invalid_charset")
var UsernameInUseCode = *errpkg.NewErrorCode("username_in_use")
func UsernameLength(message string) errpkg.StructuredError { return fieldError(UsernameLengthCode, message, usernameField) }
func UsernameFormat(message string) errpkg.StructuredError { return fieldError(UsernameFormatCode, message, usernameField) }
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

// Session Errors
var SessionInvalidLoginDetailsCode = *errpkg.NewErrorCode("session_invalid_login_details")
var SessionAlreadyLoggedInCode = *errpkg.NewErrorCode("session_already_logged_in")
var SessionAlreadyLoggedOutCode = *errpkg.NewErrorCode("session_already_logged_out")
func SessionInvalidLoginDetails(message string) errpkg.StructuredError { return plainError(SessionAlreadyLoggedInCode, message) }
func SessionAlreadyLoggedIn(message string) errpkg.StructuredError { return plainError(SessionAlreadyLoggedInCode, message) }
func SessionAlreadyLoggedOut(message string) errpkg.StructuredError { return plainError(SessionAlreadyLoggedOutCode, message) }

// Auth Errors
var AuthInvalidTokenCode = *errpkg.NewErrorCode("auth_invalid_token")
var AuthUnauthorisedCode = *errpkg.NewErrorCode("auth_unauthorised")
func AuthInvalidToken(message string) errpkg.StructuredError { return plainError(AuthInvalidTokenCode, message) }
func AuthUnauthorised(message string) errpkg.StructuredError { return plainError(AuthUnauthorisedCode, message) }