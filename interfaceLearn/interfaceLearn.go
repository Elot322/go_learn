package interfaceLearn

import (
	"fmt"
	"time"
)

type Printer interface {
	Print()
}

func Process(obj Printer) {
	obj.Print()
}

type User struct {
	name     string
	age      int
	lastName string
}

type Document struct {
	name         string
	documentType string
	date         time.Time
}

// функция Print для структуры Document
func (d Document) Print() {
	fmt.Printf("Document name: %s, type: %s, date: %s \n", d.name, d.documentType, d.date)
}

// функция Print для структуры User
func (u User) Print() {
	fmt.Printf("Hi I am %s %s and I am %d years old \n", u.name, u.lastName, u.age)
}

func InterfaceLearn() {
	u := User{name: "John", age: 24, lastName: "Smith"}
	doc := Document{name: "doc.csv", documentType: "csv", date: time.Now()}
	Process(u)
	Process(doc)
}
