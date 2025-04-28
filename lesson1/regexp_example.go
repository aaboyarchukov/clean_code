package regexp_example

import (
	"regexp"
)

// old name: IsINN
// new name: FindINN
func FindINN(stringWithINN string) bool {
	// arguments for function
	// old name: inn
	// new name: stringWithINN

	// old name: regINN_10
	// new name: regexpForINN
	regexpForINN, _ := regexp.Compile(`(\d{10})`)

	// old name: matchesINN_10
	// new name: matchesForINN
	matchesForINN := regexpForINN.FindAllString(stringWithINN, -1)

	return len(matchesForINN) != 0
}

// old name: IsPhone
// new name: FindPhone
func FindPhone(stringWithPhone string) bool {
	// arguments for function
	// old name: row
	// new name: stringWithPhone

	// old name: regPhone
	// new name: regexpForPhone
	regexpForPhone, _ := regexp.Compile(`(?:\+|\d)[\d\-\(\) ]{9,}\d`)
	return regexpForPhone.FindString(stringWithPhone) != ""
}

// old name: IsEmail
// new name: FindEmail
func FindEmail(stringWithEmail string) bool {
	// arguments for function
	// old name: row
	// new name: stringWithEmail

	// old name: regMail
	// new name: regexpForEmail
	regexpForEmail, _ := regexp.Compile(`\b[\.0-9a-zA-Z]+@[a-zA-Z\.]+\b`)

	return regexpForEmail.FindString(stringWithEmail) != ""
}
