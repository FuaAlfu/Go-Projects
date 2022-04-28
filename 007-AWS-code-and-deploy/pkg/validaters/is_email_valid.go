package validaters

import "regexp"
//func test() {}

func IsEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-z0-9.!#3$%&']")

	if len(email) < 3 || len(email) > 254 || !regexp.MatchString(email){
		return false
	}
	return true
}