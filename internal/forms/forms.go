package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//Creates a custon form struct and embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}
//Valid returns true if there are no errors
func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}

//New initializes aform structure
func New(data url.Values ) * Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
//Checks for required fields
func (f *Form) Required(fields ...string){
	for _, field := range fields{
		value := f.Get(field)
		if strings.TrimSpace(value)== ""{
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//Has chacks if form field is in post and not empty
func (f *Form) Has(field string)bool{
	x := f.Get(field)

	if x == "" {
		f.Errors.Add(field,"this field cannot be blank")
		return false
	}
	return true
}

//MiLength chacks for characters minimum length
func (f *Form) MinLength(field string, length int ) bool{
	x := f.Get(field)
	if len(x)<length{
		f.Errors.Add(field, fmt.Sprintf("this field must be atleast %d characters long", length))
		return false
	}
	return true
}
//Checks validity of email address
func (f*Form) IsEmail(field string){
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field,"Invalid email Address")
	}
}