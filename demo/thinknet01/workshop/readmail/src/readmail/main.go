package main

import (
	"fmt"
	"readmail/dao"
	"readmail/parser"
)

func main() {
	fmt.Println("Start read mail")
	m := parser.ReadFromFile("sample/1489731215.Vfd00I80501M718280.jobthai-mail01")
	fmt.Println(m.From, m.To, m.Subject)
	dao.Save(parser.Mail{})
}
