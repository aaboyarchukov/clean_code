package regexp_example

import (
	"regexp"
)

func isINN(inn string) bool {
	regINN_10, _ := regexp.Compile(`(\d{10})`)
	matchesINN_10 := regINN_10.FindAllString(inn, -1)

	return len(matchesINN_10) != 0
}

func isPhone(row string) bool {
	regPhone, _ := regexp.Compile(`(?:\+|\d)[\d\-\(\) ]{9,}\d`)
	return regPhone.FindString(row) != ""
}

func isEmail(row string) bool {
	regMail, _ := regexp.Compile(`\b[\.0-9a-zA-Z]+@[a-zA-Z\.]+\b`)

	return regMail.FindString(row) != ""
}
