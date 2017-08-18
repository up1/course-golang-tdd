package parser

import (
	"github.com/jhillyerd/enmime"
	"os"
)

type Mail struct {
	From    string
	To      string
	Subject string
}

func ReadFromFile(path string) Mail {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	env, err := enmime.ReadEnvelope(file)
	if err != nil {
		panic(err)
	}

	m := Mail{}
	m.From = env.GetHeader("From")
	alist, _ := env.AddressList("To")
	for _, addr := range alist {
		// fmt.Printf("To: %s <%s>\n", addr.Name, addr.Address)
		m.To = addr.Address
	}
	m.Subject = env.GetHeader("Subject")
	return m
}
