package password

import (
	"fmt"
	"regexp"
)

// New to new a password validator service
/**
 * input parameters:
 *    minLength, maxLength: defines password string length requirement
 *    bool parameters:
 *      requireUppercase: defines that at least contain with one uppercase letter or not
 *      requireLowercase: defines that at least contain with one lowercase letter or not
 *      requireNumber: defines that at least contain with one number or not
 *      allowSequence: defines that it is allowed sequence characters or not
 *      requireSymbol: defines that at least contain with one special symbol or not
 */
func New(minLength, maxLength uint8, requireUppercase, requireLowercase, requireNumber, allowSequence, requireSymbol bool) Service {
	return &impl{
		minLength:     minLength,
		maxLength:     maxLength,
		uppercase:     requireUppercase,
		lowercase:     requireLowercase,
		number:        requireNumber,
		sequence:      allowSequence,
		specialSymbol: requireSymbol,
	}
}

type impl struct {
	minLength     uint8
	maxLength     uint8
	uppercase     bool
	lowercase     bool
	number        bool
	sequence      bool
	specialSymbol bool
}

var (
	numRegex       = "[0-9]{1}"
	lowercaseRegex = "[a-z]{1}"
	uppercaseRegex = "[A-Z]{1}"
	symbolRegex    = "[-^&*()~_]{1}"

	// ErrFormat describe input string does not match require string length
	ErrFormat = fmt.Errorf("password format incorrect")
)

func (im *impl) Validator(pwd string) (length, uppercase, lowercase, number, symbol, sequence bool, err error) {
	lengthValid := true
	// check password length rule
	passwordLegth := uint8(len(pwd))
	if passwordLegth < im.minLength || passwordLegth > im.maxLength {
		lengthValid = false
		err = ErrFormat
	}

	// check uppercase character rule
	uppercaseValid := im.hasUppwercase(pwd)
	if !uppercaseValid {
		err = ErrFormat
	}
	// check lowercase character rule
	lowercaseValid := im.hasLowercase(pwd)
	if !lowercaseValid {
		err = ErrFormat
	}

	// check special symbol character rule
	symbolValid := im.hasSymbol(pwd)
	if !symbolValid {
		err = ErrFormat
	}

	// check number character rule
	numberValid := im.hasNumber(pwd)
	if !numberValid {
		err = ErrFormat
	}

	sequenceValid := im.hasSequence(pwd)
	if !sequenceValid {
		err = ErrFormat
	}

	return lengthValid, uppercaseValid, lowercaseValid, numberValid, symbolValid, sequenceValid, err
}

func (im *impl) hasNumber(input string) bool {
	check := regexp.MustCompile(numRegex).MatchString(input)
	required := im.number

	if check != required {
		return false
	}

	return true
}

func (im *impl) hasUppwercase(input string) bool {
	check := regexp.MustCompile(uppercaseRegex).MatchString(input)
	required := im.uppercase

	if check != required {
		return false
	}

	return true
}

func (im *impl) hasLowercase(input string) bool {
	check := regexp.MustCompile(lowercaseRegex).MatchString(input)
	required := im.lowercase

	if check != required {
		return false
	}

	return true
}

func (im *impl) hasSymbol(input string) bool {
	check := regexp.MustCompile(symbolRegex).MatchString(input)
	required := im.specialSymbol

	if check != required {
		return false
	}

	return true
}

func (im *impl) hasSequence(input string) bool {
	byteArr := []byte(input)
	allowed := im.sequence
	hasSequenceChar := false

	for idx := range byteArr {
		if idx > 0 {
			if int(byteArr[idx])-int(byteArr[idx-1]) == 1 || int(byteArr[idx])-int(byteArr[idx-1]) == -1 {
				hasSequenceChar = true
			}
		}
	}

	if !allowed && hasSequenceChar {
		return false
	}

	return true
}
