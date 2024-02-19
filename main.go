package main

import "errors"

type Password string

func (p Password) Validate() error {
	if len(p) < 6 {
		return errors.New("password too short")
	}
	return nil
}

func main() {
	password := Password("12")
	err := password.Validate()
	if err != nil {
		panic(err)
	}
}
