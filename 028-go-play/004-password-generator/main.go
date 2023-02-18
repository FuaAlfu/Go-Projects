package main

import (
"crypto/rand"
"fmt"

)

/*
generateRandomBytes returns securely generated random bytes.
It will return an error if the system's secure random number generator failed to function correctly, in which case the caller should not continue.
*/
func generateRandomBytes(n int) ([]byte, error) {
b := make([]byte, n)
_, err := rand.Read(b)

// Note that err == nil only if we read len(b) //bytes.
if err != nil {
return nil, err
}

return b, nil
}

/* GenerateString returns a securely generated random string. It will return an error if the system's secure random number generator fails to function correctly, in which case the caller should not continue. 25 is the length of the password being returned by this function. //This code to generate passwords modified from https://golangcode.com/generate-random-string/
*/
func GenerateString(n int) (string, error) {
const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!?#@$&*&"
//mixed of uppercase charcters & //lowercase characters and provides //additional chars as option making it harder //for hackers

bytes, err := generateRandomBytes(n)

for i, b := range bytes { //loops through each byte stored in 'bytes'
bytes[i] = alphanum[b % byte(len(alphanum))] //sets each index at 'byes' location to a character from alphanumeric list dependant on b value modulus len('alphanum') returning only values from 0 - 62 hence generating random output with specified shift across numbers & characters in list

}

return string(bytes), err
/*
location for array with new/modified indexed values returns index positions as passwords
*/
}

func main() {
/*
main used as demonstration of password generation functions
*/
x,err := GenerateString(257) //defines x equal to returned string value of GenerateString() with the parameter of 25 denoting length of output
if err != nil {
	panic(err)
}
fmt.Println("Your password is: ", x ) //returns generated password

}