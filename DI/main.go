package main

import "fmt"

type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDB struct {}

func (db DefaultDB) GetUser() string {
	return "bob"
}

func (db DefaultDB) GetAllUsers() []string {
	return []string{"Bob", "Tom", "John"}
}

type FamouseDB struct {}

func (db FamouseDB) GetUser() string{
	return "Sam"
}

func (db FamouseDB) GetAllUsers() []string {
	return []string{"Sam", "Bill", "Roy"}
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct {}

func (ng NiceGreeter) Greet(userName string) {
	fmt.Printf("Hello %s", userName)
}

type Program struct {
	Db Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	//db := DefaultDB{}
	db := FamouseDB{}
	greeter := NiceGreeter{}
	p := Program{
		Db: db,
		Greeter: greeter,
	}

	p.Execute()
}