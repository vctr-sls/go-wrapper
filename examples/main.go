package main

import (
	"fmt"

	vctr "github.com/vctr-sls/go-wrapper"
)

func main() {
	c := vctr.NewClient("http://localhost:5000")
	res, err := c.Auth.Login(&vctr.LoginModel{
		Ident:    "root",
		Password: "root",
		Remember: true,
	})
	fmt.Println(res, err)

	{
		res, _ := c.Users.List(100, 0)
		fmt.Println(res[0])

		fmt.Println(res[0].LinksCount())
	}
}
