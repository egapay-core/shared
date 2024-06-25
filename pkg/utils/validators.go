package utils

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"
)

var (
	supportedLanguages = []string{"en", "fr"}

	ErrNoEmailAddress      = status.Errorf(codes.InvalidArgument, "email address is required")
	ErrInvalidEmailAddress = status.Errorf(codes.InvalidArgument, "invalid email address")

	ErrNoPassword      = status.Errorf(codes.InvalidArgument, "password cannot be empty")
	ErrInvalidPassword = status.Errorf(codes.InvalidArgument, "password must be at least 8 characters long and may contain at least one special character")

	ErrNoPhoneNumber       = status.Errorf(codes.InvalidArgument, "phone number is required")
	ErrInvalidPhoneNumber  = status.Errorf(codes.InvalidArgument, "invalid phone number")
	ErrPhoneNumberDialCode = status.Errorf(codes.InvalidArgument, "country is required")

	ErrNoDialCode      = status.Errorf(codes.InvalidArgument, "dial code is required")
	ErrInvalidDialCode = status.Errorf(codes.InvalidArgument, "invalid dial code")

	ErrNoCountryCode      = status.Errorf(codes.InvalidArgument, "country code is required")
	ErrInvalidCountryCode = status.Errorf(codes.InvalidArgument, "invalid country code")

	ErrNoLanguageId      = status.Errorf(codes.InvalidArgument, "language id is required")
	ErrInvalidLanguageId = status.Errorf(codes.InvalidArgument, "unsupported language id")
)

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return ErrNoEmailAddress
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if ok, _ := regexp.MatchString(emailRegex, email); !ok {
		return ErrInvalidEmailAddress
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordRegex := `^[a-zA-Z0-9!@#$&()\-.+]{8,}$`
	if len(password) == 0 {
		return ErrNoPassword
	}
	if ok, _ := regexp.MatchString(passwordRegex, password); !ok {
		return ErrInvalidPassword
	}

	return nil
}

func ValidatePhoneNumber(phoneNumber, countryCode string) error {
	if len(phoneNumber) == 0 {
		return ErrNoPhoneNumber
	}

	if len(countryCode) == 0 {
		if strings.HasPrefix(phoneNumber, "0") {
			return ErrPhoneNumberDialCode
		}

		phoneNumberWithDialCodeRegex := `^\+[0-9]{10,15}$`
		if ok, _ := regexp.MatchString(phoneNumberWithDialCodeRegex, phoneNumber); !ok {
			return ErrInvalidPhoneNumber
		}
		return nil
	}

	// if dial code is provided, reformat the phone number to include the dial code
	if len(countryCode) > 0 && strings.HasPrefix(phoneNumber, "0") {
		phoneNumber = fmt.Sprintf("%s%s", countryCode, strings.TrimLeft(phoneNumber, "0"))
	}

	// valid phone numbers: 0241234567, 233241234567, +233241234567
	phoneNumberWithDialCodeRegex := fmt.Sprintf(`^(\%s|0)[0-9]{9,12}$`, countryCode)
	if ok, _ := regexp.MatchString(phoneNumberWithDialCodeRegex, phoneNumber); !ok {
		return ErrInvalidPhoneNumber
	}

	return nil
}

func ValidateDialCode(code string) error {
	if len(code) == 0 {
		return ErrNoDialCode
	}

	dialCodeRegex := `^\+[0-9]{1,3}$`
	if ok, _ := regexp.MatchString(dialCodeRegex, code); !ok {
		return ErrInvalidDialCode
	}

	return nil
}

func ValidateCountryCode(code string) error {
	if len(code) == 0 {
		return ErrNoCountryCode
	}

	dialCodeRegex := `^[A-Z0-9]{6}$`
	if ok, _ := regexp.MatchString(dialCodeRegex, code); !ok {
		return ErrInvalidCountryCode
	}

	return nil

}

func ValidateLanguageId(languageId string) error {
	if len(languageId) == 0 {
		return ErrNoLanguageId
	}

	languageIdRegex := `^[a-z]{2,3}$`
	if ok, _ := regexp.MatchString(languageIdRegex, languageId); !ok {
		return ErrInvalidLanguageId
	}

	for _, lang := range supportedLanguages {
		if strings.HasPrefix(languageId, lang) {
			return nil
		}
	}

	return ErrInvalidLanguageId
}
