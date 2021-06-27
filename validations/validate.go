package validations

import (
	"cust-service/model"
	"net/mail"
	"net/url"
	"regexp"
)

var regexpUname = regexp.MustCompile("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+){0,2}$")

var regexPass = regexp.MustCompile("^?=.*?[A-Z]?=.*?[#?!@$%^&*-].{8,}$")

//user validation
func ValidUser(u model.User) url.Values {

	errs := url.Values{}

	if u.UserName == "" && len(u.UserName) <= 30 {
		errs.Add("username", "The username is required!")
	}

	if !regexpUname.MatchString(u.UserName) {
		errs.Add("username", "Username may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen.!")
	}

	if !regexPass.MatchString(u.Password) {
		errs.Add("Password", "Minimum 8 Character, 1 capital letter and 1 Special character!")
	}

	return errs
}

//item validation
func ValidItem(i model.Item) url.Values {

	errs := url.Values{}

	if len(i.ItemName) == 0 || CountDigits(int(i.Qty)) <= 0 || CountDigits(int(i.Rate)) <= 0 {
		errs.Add("ItemName, rate qty", "itemname, Rate, Qty should not be blank or Qty and Rate can't be zero or negative value!")
	}

	if len(i.ItemName) > 20 {
		errs.Add("ItemName", "item name should not be grater than 20 character!")
	}

	return errs
}

//customer validation
func ValidCustomer(c model.Customer) url.Values {

	errs := url.Values{}

	if len(c.FirstName) == 0 || len(c.FirstName) > 30 || len(c.LastName) == 0 || len(c.LastName) > 30 || CountDigits(int(c.Phone)) == 0 || len(c.Email) == 0 {
		errs.Add("Firstname, Lastname, Phone, Email", "all fields are mandatory and FirstName & LastName cant be more than 30 characters!")
	}

	if CountDigits(int(c.Phone)) < 10 || CountDigits(int(c.Phone)) > 10 {
		errs.Add("Phone", "phone number should not be greater and less than 10 !")
	}

	if len(c.FirstName) > 30 || len(c.LastName) > 30 {
		errs.Add("firstname and lastname", "firstname and lastname should be less than  30 character")
	}

	if _, err := mail.ParseAddress(c.Email); err != nil {
		errs.Add("email", "please enter valid email address")
	}
	return errs
}

func CountDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}
