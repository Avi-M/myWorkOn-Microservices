package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person{
		UserName: "Avinash maurya",
		email:    "avinashup45@gmail.com",
	}
	t.Execute(os.Stdout, p)
}
