package main

import (
	"fmt"

	vctr "github.com/vctr-sls/go-wrapper"
)

func main() {
	c := vctr.NewClient("http://localhost:5000")
	res, err := c.Login(&vctr.LoginModel{
		Ident:    "root",
		Password: "root",
		Remember: true,
	})
	fmt.Println(res, err)

	{
		res, err := c.Users.GetMe()
		fmt.Println(res, err)
	}
}
