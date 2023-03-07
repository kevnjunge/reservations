package forms

import (
	"net/http"
	"net/url"
)

//Creates a custon form struct and embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

//New initializes aform structure
func New(data url.Values ) * Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has chacks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request)bool{
	x := r.Form.Get(field)

	if x == "" {
		return false
	}
	return true
}