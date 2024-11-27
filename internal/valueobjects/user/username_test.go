package user

import (
	"testing"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

func TestNewUsername(t *testing.T) {
	tests := []struct {
		input    string
		expected []errpkg.ErrorCode
	}{
		// Length Checks
		// - must be between 8-24 characters long (inclusive)
		{"username", nil},
		{"usernameusername", nil},
		{"usernameusernameusername", nil},
		{"short", []errpkg.ErrorCode {errs.UsernameLengthCode}},
		{"usernametoolonggggggggggg", []errpkg.ErrorCode {errs.UsernameLengthCode}},

		// Valid Charset Checks
		// - must only contain alphanumeric, underscore and dash
		{"user-name", nil},
		{"_username_", nil},
		{"_u-s-e-r_", nil},
		{"usernamé",  []errpkg.ErrorCode {errs.UsernameCharsetCode}},
		{"username!", []errpkg.ErrorCode {errs.UsernameCharsetCode}},
		{"username😃", []errpkg.ErrorCode {errs.UsernameCharsetCode}},
		{"        ", []errpkg.ErrorCode {errs.UsernameCharsetCode}},

		// Valid Format Checks
		// - cannot have more than 1 allowed special character in a row
		{"user--name", []errpkg.ErrorCode {errs.UsernameFormatCode}},
		{"__username__", []errpkg.ErrorCode {errs.UsernameFormatCode}},
		{"user-_-name", []errpkg.ErrorCode {errs.UsernameFormatCode}},

		// Restricted Word Checks
		// - cannot contain restricted or reserved keywords
		{"username_admiral", nil},
		{"username_truck", nil},
		{"username_admin", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		{"username_FuCk", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		{"username_c_u_n_t", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		{"username_biiitch", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		{"username_5lut", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		{"username_assh_le", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},
		// {"f_ckkaiiidon", []errpkg.ErrorCode {errs.UsernameRestrictedCode}},

		// Multiple Error Checks
		{"!!!!!!", []errpkg.ErrorCode {errs.UsernameLengthCode, errs.UsernameCharsetCode}},
		{"!!!__!!!", []errpkg.ErrorCode {errs.UsernameFormatCode, errs.UsernameCharsetCode}},
		{"admin", []errpkg.ErrorCode {errs.UsernameLengthCode}},
		{"fu__ck!", []errpkg.ErrorCode {errs.UsernameLengthCode, errs.UsernameCharsetCode, errs.UsernameFormatCode}},
	}

	for _, test := range tests {
		httpStack := errpkg.NewHttpStack()
		result := NewUsername(test.input, httpStack)

		// Check For Nil
		if (test.expected == nil) {
			if !httpStack.IsEmpty() {
				t.Errorf("NewUsername(%s) = %v; expected no errors; recieved: %v", test.input, result, httpStack.Errors)
			}

			continue
		}

		// Otherwise Check Errors
		for _, code := range test.expected {
			if !httpStack.Contains(code) {
				t.Errorf("NewUsername(%s) = %v; expected error code %v", test.input, result, code)
			}
		}
	}
}