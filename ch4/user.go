package main

import "errors"

type PhoneNumber struct {
	countryCode int
	number      [10]int
}
type User struct {
	name        string
	phoneNumber PhoneNumber
}

func GetUsersMap(names []string, phoneNumbers []PhoneNumber) (map[string]User, error) {
	n := len(names)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("number of names and phone numbers differ")
	} else {
		Users := make(map[string]User)
		for i := 0; i < n; i++ {
			Users[names[i]] = User{
				name:        names[i],
				phoneNumber: phoneNumbers[i],
			}
		}
		return Users, nil
	}
}
func deleteUsers(users map[string]User, name string) (deleted bool, err error) {
	//check if the user exists
	if _, b := users[name]; b == false {
		return false, errors.New("user to de deleted, does not exist")
	} else {
		delete(users, name)
		return true, nil
	}
}
