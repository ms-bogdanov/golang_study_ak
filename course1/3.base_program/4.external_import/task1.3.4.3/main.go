package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	name := fake.FullName()
	email := fake.EmailAddress()
	phone := fake.Phone()
	address := fake.StreetAddress()
	return fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\nEmail: %s", name, address, phone, email)

}

func main() {
	fmt.Println(GenerateFakeData())
}
