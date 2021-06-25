package validations

import (
	"cust-service/model"
	"net/url"
	"regexp"
)

var regexpUname = regexp.MustCompile("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+){0,2}$")

var regexPass = regexp.MustCompile("^?=.*?[A-Z]?=.*?[#?!@$%^&*-].{8,}$")

func ValidUser(u model.User) url.Values {

	errs := url.Values{}

	if u.UserName == "" && len(u.UserName) <= 30 {
		errs.Add("username", "The username is required!")
	}

	if !regexpUname.MatchString(u.UserName) {
		errs.Add("username", "Username may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen.!")
	}

	if !regexPass.MatchString(u.Password) {
		errs.Add("Password", "(Minimum 8 Character, 1 capital letter and 1 Special character)!")
	}

	return errs
}
