package utils

import "regexp"

var (
	CheckEmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	CheckPhoneRX = regexp.MustCompile(`^1[3-9]\d{9}`)
)

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) String() string {
	if len(v.Errors) == 0 {
		return ""
	}
	var msg string
	for _, err := range v.Errors {
		msg += err + ","
	}
	msg = msg[:len(msg)-2]
	return msg
}

func (v *Validator) AddError(key, msg string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = msg
	}
}

func (v *Validator) CheckEmail(value string, key string, msg string) {
	if !CheckEmailRX.MatchString(value) {
		v.AddError(key, msg)
	}
}

func (v *Validator) CheckPhone(value string, key string, msg string) {
	if !CheckPhoneRX.MatchString(value) {
		v.AddError(key, msg)
	}
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}
