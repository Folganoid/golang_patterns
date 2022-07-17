package main

import "fmt"

type structHere struct {
	N1, N2 int
}

type InterfaceHere interface {
	Sum() int
}

type otherStruct struct {
	A, B int
}

func (s otherStruct) Sum() int {
	return s.A + s.B
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type user struct {
	FIO, Address, Phone string
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
	Show()
}

func main() {

	var a InterfaceHere
	sh := structHere{1,2}
	os := otherStruct{2,4}

	a = &sh
	fmt.Println(a.Sum())
	a = &os
	fmt.Println(a.Sum())
	a = os
	fmt.Println(a.Sum())

	var i InterfaceHere = otherStruct{9, 5}
	fmt.Println(i.Sum())

	user := NewUser("Andrey", "Sumy", "104")
	user.Show()

}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

func (u *user) ChangeFIO(newFIO string) {
	u.FIO = newFIO
}

func (u *user) Show() {
	fmt.Print(u.FIO, "\n")
	fmt.Print(u.Address, "\n")
	fmt.Print(u.Phone, "\n")
}

//realisation
func NewUser(fio, address, phone string) User {
	u := user {
		FIO: fio,
		Address: address,
		Phone: phone,
	}

	return &u
}